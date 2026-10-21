package handler

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"skillshub/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

const testStateCookie = "oauth_state"

func findCookie(w *httptest.ResponseRecorder, name string) *http.Cookie {
	for _, c := range w.Result().Cookies() {
		if c.Name == name {
			return c
		}
	}
	return nil
}

// newAuthEngine wires only the OAuth routes and points the handler at a fake
// Google served by httptest, so no real network access is needed.
func newAuthEngine(t *testing.T, extraEnv map[string]string) (*gin.Engine, *AuthHandler, *config.Config, *int32) {
	t.Helper()
	db := testDB(t)
	cfg := testConfig(t, extraEnv)
	h := NewAuthHandler(db, cfg)

	var tokenHits int32
	google := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			atomic.AddInt32(&tokenHits, 1)
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"at","token_type":"Bearer","expires_in":3600}`))
		case "/userinfo":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"g-123","email":"new@example.com","name":"New User","picture":"http://x/p.png"}`))
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(google.Close)
	h.oauthConfig.Endpoint = oauth2.Endpoint{AuthURL: google.URL + "/auth", TokenURL: google.URL + "/token"}
	h.userInfoURL = google.URL + "/userinfo"

	r := gin.New()
	r.GET("/api/auth/google", h.GoogleLogin)
	r.GET("/api/auth/google/callback", h.GoogleCallback)
	return r, h, cfg, &tokenHits
}

func startLogin(t *testing.T, r http.Handler) (state string, cookie *http.Cookie) {
	t.Helper()
	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodGet, "/api/auth/google", nil))
	if w.Code != http.StatusOK {
		t.Fatalf("login status %d", w.Code)
	}
	var body struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	i := strings.Index(body.URL, "state=")
	if i < 0 {
		t.Fatalf("no state in url %q", body.URL)
	}
	state = body.URL[i+len("state="):]
	if j := strings.IndexByte(state, '&'); j >= 0 {
		state = state[:j]
	}
	cookie = findCookie(w, testStateCookie)
	if cookie == nil {
		t.Fatal("login did not set the state cookie")
	}
	return state, cookie
}

func TestGoogleLoginSetsRandomStateCookie(t *testing.T) {
	r, _, _, _ := newAuthEngine(t, nil)
	seen := map[string]bool{}
	for i := 0; i < 5; i++ {
		state, c := startLogin(t, r)
		if state != c.Value {
			t.Fatalf("authorize URL state %q != cookie %q", state, c.Value)
		}
		raw, err := base64.RawURLEncoding.DecodeString(c.Value)
		if err != nil || len(raw) != 32 {
			t.Fatalf("state must be 32 random bytes base64url (len=%d err=%v)", len(raw), err)
		}
		if seen[c.Value] {
			t.Fatal("state repeated")
		}
		seen[c.Value] = true
		if !c.HttpOnly {
			t.Error("cookie must be HttpOnly")
		}
		if c.SameSite != http.SameSiteLaxMode {
			t.Errorf("SameSite = %v, want Lax", c.SameSite)
		}
		if c.Path != "/api/auth" {
			t.Errorf("Path = %q", c.Path)
		}
		if c.MaxAge != 600 {
			t.Errorf("MaxAge = %d", c.MaxAge)
		}
		if c.Secure {
			t.Error("Secure must be off for an http:// frontend")
		}
	}
}

func TestGoogleLoginCookieSecureForHTTPSFrontend(t *testing.T) {
	r, _, _, _ := newAuthEngine(t, map[string]string{"FRONTEND_URL": "https://skills.example.com"})
	_, c := startLogin(t, r)
	if !c.Secure {
		t.Fatal("cookie must be Secure for an https:// frontend")
	}
}

func TestGoogleCallbackRejectsInvalidState(t *testing.T) {
	r, _, _, tokenHits := newAuthEngine(t, nil)
	good, goodCookie := startLogin(t, r)

	cases := []struct {
		name   string
		query  string
		cookie string // empty = no cookie
	}{
		{"no cookie, no state", "code=abc", ""},
		{"cookie but no state param", "code=abc", goodCookie.Value},
		{"state param but no cookie", "code=abc&state=" + good, ""},
		{"mismatch", "code=abc&state=" + good + "x", goodCookie.Value},
		{"different value", "code=abc&state=AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA", goodCookie.Value},
		{"empty state and cookie", "code=abc&state=", ""},
		{"prefix of cookie", "code=abc&state=" + good[:10], goodCookie.Value},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/auth/google/callback?"+tc.query, nil)
			if tc.cookie != "" {
				req.AddCookie(&http.Cookie{Name: testStateCookie, Value: tc.cookie})
			}
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want 400", w.Code)
			}
			if loc := w.Header().Get("Location"); loc != "" {
				t.Fatalf("unexpected redirect %q", loc)
			}
			c := findCookie(w, testStateCookie)
			if c == nil || c.MaxAge >= 0 || c.Value != "" {
				t.Fatalf("state cookie must be cleared on rejection, got %+v", c)
			}
		})
	}
	if n := atomic.LoadInt32(tokenHits); n != 0 {
		t.Fatalf("code exchange attempted %d times despite invalid state", n)
	}
}

func TestGoogleCallbackSuccessUsesFragmentAndClearsState(t *testing.T) {
	r, _, cfg, tokenHits := newAuthEngine(t, nil)
	state, c := startLogin(t, r)

	req := httptest.NewRequest(http.MethodGet, "/api/auth/google/callback?code=abc&state="+state, nil)
	req.AddCookie(&http.Cookie{Name: testStateCookie, Value: c.Value})
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusTemporaryRedirect && w.Code != http.StatusFound {
		t.Fatalf("status = %d body=%s", w.Code, w.Body.String())
	}
	if atomic.LoadInt32(tokenHits) != 1 {
		t.Fatalf("expected exactly one code exchange")
	}
	loc := w.Header().Get("Location")
	prefix := cfg.FrontendURL + "/auth/callback#token="
	if !strings.HasPrefix(loc, prefix) {
		t.Fatalf("Location = %q, want prefix %q", loc, prefix)
	}
	if strings.Contains(loc, "?") {
		t.Fatalf("token must not travel in the query string: %q", loc)
	}
	tokenStr := strings.TrimPrefix(loc, prefix)
	parsed, err := jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) { return []byte(cfg.JWTSecret), nil },
		jwt.WithValidMethods([]string{"HS256"}), jwt.WithExpirationRequired())
	if err != nil || !parsed.Valid {
		t.Fatalf("issued token invalid: %v", err)
	}
	if got := parsed.Claims.(jwt.MapClaims)["email"]; got != "new@example.com" {
		t.Fatalf("email claim = %v", got)
	}
	if cc := findCookie(w, testStateCookie); cc == nil || cc.MaxAge >= 0 {
		t.Fatalf("state cookie must be cleared after success, got %+v", cc)
	}
	if cc := w.Header().Get("Cache-Control"); !strings.Contains(cc, "no-store") {
		t.Errorf("Cache-Control = %q, want no-store", cc)
	}
}
