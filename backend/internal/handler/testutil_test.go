package handler

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"skillshub/internal/config"
	"skillshub/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const testJWTSecret = "handler-test-secret-long-enough-0123456789"

var memCounter int64

func init() { gin.SetMode(gin.TestMode) }

func testConfig(t *testing.T, extra map[string]string) *config.Config {
	t.Helper()
	env := map[string]string{
		"JWT_SECRET":   testJWTSecret,
		"FRONTEND_URL": "http://frontend.test",
	}
	for k, v := range extra {
		env[k] = v
	}
	cfg, err := config.LoadFrom(func(k string) string { return env[k] })
	if err != nil {
		t.Fatal(err)
	}
	return cfg
}

// testDB returns a fresh, seeded in-memory SQLite database with foreign keys on.
func testDB(t *testing.T) *sql.DB {
	t.Helper()
	n := atomic.AddInt64(&memCounter, 1)
	db, err := service.OpenDB(fmt.Sprintf("file:handlertest%d?mode=memory&cache=shared", n))
	if err != nil {
		t.Fatalf("open test db: %v", err)
	}
	// One connection keeps the shared in-memory database alive and makes any
	// accidental nested use of the pool show up as a hang rather than flakiness.
	db.SetMaxOpenConns(1)
	t.Cleanup(func() { db.Close() })
	return db
}

func testRouter(t *testing.T) (*gin.Engine, *sql.DB, *config.Config) {
	t.Helper()
	db := testDB(t)
	cfg := testConfig(t, nil)
	r, err := NewRouter(db, cfg)
	if err != nil {
		t.Fatal(err)
	}
	return r, db, cfg
}

func tokenFor(t *testing.T, cfg *config.Config, userID int64) string {
	t.Helper()
	s, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   "tester@example.com",
		"name":    "Tester",
		"exp":     time.Now().Add(time.Hour).Unix(),
	}).SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		t.Fatal(err)
	}
	return s
}

func doJSON(r http.Handler, method, path, token string, body interface{}) *httptest.ResponseRecorder {
	var buf bytes.Buffer
	if body != nil {
		_ = json.NewEncoder(&buf).Encode(body)
	}
	req := httptest.NewRequest(method, path, &buf)
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func count(t *testing.T, db *sql.DB, query string, args ...interface{}) int {
	t.Helper()
	var n int
	if err := db.QueryRow(query, args...).Scan(&n); err != nil {
		t.Fatal(err)
	}
	return n
}
