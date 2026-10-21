package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func clientIPFor(t *testing.T, remoteAddr, xff string) string {
	t.Helper()
	r, _, _ := testRouter(t)
	r.GET("/_ip", func(c *gin.Context) { c.String(http.StatusOK, c.ClientIP()) })
	req := httptest.NewRequest(http.MethodGet, "/_ip", nil)
	req.RemoteAddr = remoteAddr
	if xff != "" {
		req.Header.Set("X-Forwarded-For", xff)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w.Body.String()
}

func TestOnlyLocalProxyIsTrusted(t *testing.T) {
	// A direct client cannot spoof its address with X-Forwarded-For.
	if got := clientIPFor(t, "203.0.113.9:4000", "1.2.3.4"); got != "203.0.113.9" {
		t.Errorf("untrusted peer: ClientIP = %q, want 203.0.113.9", got)
	}
	// The local reverse proxy is trusted to forward the real client address.
	if got := clientIPFor(t, "127.0.0.1:4000", "198.51.100.7"); got != "198.51.100.7" {
		t.Errorf("trusted proxy: ClientIP = %q, want 198.51.100.7", got)
	}
}

func TestAPIStillServesPublicRoutes(t *testing.T) {
	r, _, _ := testRouter(t)
	for _, path := range []string{"/api/skills", "/api/skills/featured", "/api/skills/categories", "/api/skills/1", "/api/skills/1/comments", "/api/articles", "/api/articles/1"} {
		w := doJSON(r, http.MethodGet, path, "", nil)
		if w.Code != http.StatusOK {
			t.Errorf("GET %s = %d", path, w.Code)
		}
	}
	if w := doJSON(r, http.MethodGet, "/api/auth/me", "", nil); w.Code != http.StatusUnauthorized {
		t.Errorf("GET /api/auth/me anonymous = %d, want 401", w.Code)
	}
}
