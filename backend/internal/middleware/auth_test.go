package middleware

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const testSecret = "unit-test-secret-that-is-long-enough-123"

func init() { gin.SetMode(gin.TestMode) }

func newRouter(mw gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.GET("/protected", mw, func(c *gin.Context) {
		uid, _ := c.Get("user_id")
		email, _ := c.Get("email")
		name, _ := c.Get("name")
		c.JSON(http.StatusOK, gin.H{"user_id": uid, "email": email, "name": name})
	})
	return r
}

func do(r http.Handler, header string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	if header != "" {
		req.Header.Set("Authorization", header)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func sign(t *testing.T, method jwt.SigningMethod, key interface{}, claims jwt.MapClaims) string {
	t.Helper()
	s, err := jwt.NewWithClaims(method, claims).SignedString(key)
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func goodClaims() jwt.MapClaims {
	return jwt.MapClaims{
		"user_id": 42,
		"email":   "a@example.com",
		"name":    "Alice",
		"exp":     time.Now().Add(time.Hour).Unix(),
	}
}

func hs256(t *testing.T, claims jwt.MapClaims) string {
	return sign(t, jwt.SigningMethodHS256, []byte(testSecret), claims)
}

func TestAuthRequiredAcceptsValidToken(t *testing.T) {
	r := newRouter(AuthRequired(testSecret))
	w := do(r, "Bearer "+hs256(t, goodClaims()))
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d body=%s", w.Code, w.Body.String())
	}
	var body map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	if body["user_id"] != float64(42) || body["email"] != "a@example.com" || body["name"] != "Alice" {
		t.Fatalf("unexpected context values: %v", body)
	}
}

func TestAuthRequiredRejectsBadTokens(t *testing.T) {
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}
	valid := hs256(t, goodClaims())

	expired := goodClaims()
	expired["exp"] = time.Now().Add(-time.Hour).Unix()

	noExp := goodClaims()
	delete(noExp, "exp")

	missingUser := goodClaims()
	delete(missingUser, "user_id")
	missingEmail := goodClaims()
	delete(missingEmail, "email")
	missingName := goodClaims()
	delete(missingName, "name")

	wrongUserType := goodClaims()
	wrongUserType["user_id"] = "42"
	wrongEmailType := goodClaims()
	wrongEmailType["email"] = 12345
	wrongNameType := goodClaims()
	wrongNameType["name"] = []string{"x"}

	parts := strings.Split(valid, ".")
	tampered := parts[0] + "." + parts[1] + "." + strings.Repeat("A", len(parts[2]))
	// change payload but keep the original signature
	forgedPayload := hs256(t, func() jwt.MapClaims { c := goodClaims(); c["user_id"] = 1; return c }())
	fp := strings.Split(forgedPayload, ".")
	swapped := fp[0] + "." + fp[1] + "." + parts[2]

	cases := map[string]string{
		"no header":        "",
		"empty bearer":     "Bearer ",
		"garbage":          "Bearer not-a-jwt",
		"expired":          "Bearer " + hs256(t, expired),
		"no exp":           "Bearer " + hs256(t, noExp),
		"tampered sig":     "Bearer " + tampered,
		"swapped payload":  "Bearer " + swapped,
		"wrong secret":     "Bearer " + sign(t, jwt.SigningMethodHS256, []byte("another-secret-another-secret-another"), goodClaims()),
		"alg none":         "Bearer " + sign(t, jwt.SigningMethodNone, jwt.UnsafeAllowNoneSignatureType, goodClaims()),
		"alg RS256":        "Bearer " + sign(t, jwt.SigningMethodRS256, rsaKey, goodClaims()),
		"alg HS384":        "Bearer " + sign(t, jwt.SigningMethodHS384, []byte(testSecret), goodClaims()),
		"missing user_id":  "Bearer " + hs256(t, missingUser),
		"missing email":    "Bearer " + hs256(t, missingEmail),
		"missing name":     "Bearer " + hs256(t, missingName),
		"user_id string":   "Bearer " + hs256(t, wrongUserType),
		"email not string": "Bearer " + hs256(t, wrongEmailType),
		"name not string":  "Bearer " + hs256(t, wrongNameType),
	}
	r := newRouter(AuthRequired(testSecret))
	for name, header := range cases {
		t.Run(name, func(t *testing.T) {
			w := do(r, header)
			if w.Code != http.StatusUnauthorized {
				t.Fatalf("status = %d, want 401 (body=%s)", w.Code, w.Body.String())
			}
		})
	}
}

func TestAuthOptional(t *testing.T) {
	r := newRouter(AuthOptional(testSecret))

	// anonymous passes through with no identity
	w := do(r, "")
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"user_id":null`) {
		t.Fatalf("anonymous: %d %s", w.Code, w.Body.String())
	}
	// valid token sets identity
	w = do(r, "Bearer "+hs256(t, goodClaims()))
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"user_id":42`) {
		t.Fatalf("valid: %d %s", w.Code, w.Body.String())
	}
	// malformed claims must not panic and must not set identity
	bad := goodClaims()
	bad["email"] = 7
	w = do(r, "Bearer "+hs256(t, bad))
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"user_id":null`) {
		t.Fatalf("bad claims: %d %s", w.Code, w.Body.String())
	}
	// alg none must not authenticate
	none := sign(t, jwt.SigningMethodNone, jwt.UnsafeAllowNoneSignatureType, goodClaims())
	w = do(r, "Bearer "+none)
	if !strings.Contains(w.Body.String(), `"user_id":null`) {
		t.Fatalf("alg none authenticated: %s", w.Body.String())
	}
}

func TestMaxBodyBytes(t *testing.T) {
	r := gin.New()
	r.Use(MaxBodyBytes(1024))
	r.POST("/echo", func(c *gin.Context) {
		b, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatus(http.StatusRequestEntityTooLarge)
			return
		}
		c.String(http.StatusOK, "%d", len(b))
	})

	small := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/echo", bytes.NewReader(make([]byte, 1024)))
	r.ServeHTTP(small, req)
	if small.Code != http.StatusOK {
		t.Fatalf("1024 bytes should pass, got %d", small.Code)
	}

	// declared Content-Length too large
	big := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/echo", bytes.NewReader(make([]byte, 2048)))
	r.ServeHTTP(big, req)
	if big.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("declared oversize: got %d, want 413", big.Code)
	}

	// unknown length (chunked) must still be capped while reading
	chunked := httptest.NewRecorder()
	req = httptest.NewRequest(http.MethodPost, "/echo", bytes.NewReader(make([]byte, 4096)))
	req.ContentLength = -1
	r.ServeHTTP(chunked, req)
	if chunked.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("chunked oversize: got %d, want 413", chunked.Code)
	}
}
