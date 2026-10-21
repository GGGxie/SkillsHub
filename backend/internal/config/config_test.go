package config

import (
	"strings"
	"testing"
)

const validSecret = "0123456789abcdef0123456789abcdef-ok" // 35 chars

func env(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestLoadFromRejectsBadJWTSecret(t *testing.T) {
	cases := map[string]string{
		"missing":     "",
		"legacy":      "your-secret-key-change-in-production",
		"too short":   "short-secret",
		"31 chars":    strings.Repeat("a", 31),
		"whitespaces": strings.Repeat(" ", 40),
	}
	for name, secret := range cases {
		t.Run(name, func(t *testing.T) {
			cfg, err := LoadFrom(env(map[string]string{"JWT_SECRET": secret}))
			if err == nil {
				t.Fatalf("expected error for secret %q, got cfg %+v", name, cfg)
			}
			if cfg != nil {
				t.Fatalf("expected nil config on error")
			}
			if secret != "" && strings.TrimSpace(secret) != "" && strings.Contains(err.Error(), secret) {
				t.Fatalf("error message must not echo the secret")
			}
		})
	}
}

func TestLoadFromAcceptsStrongSecret(t *testing.T) {
	// 43 chars, same length as the production secret
	secret := strings.Repeat("k", 43)
	cfg, err := LoadFrom(env(map[string]string{"JWT_SECRET": secret}))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.JWTSecret != secret {
		t.Fatalf("secret not propagated")
	}
	boundary := strings.Repeat("b", 32)
	if _, err := LoadFrom(env(map[string]string{"JWT_SECRET": boundary})); err != nil {
		t.Fatalf("32-char secret should be accepted: %v", err)
	}
}

func TestLoadFromDefaults(t *testing.T) {
	cfg, err := LoadFrom(env(map[string]string{"JWT_SECRET": validSecret}))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Port != "8080" {
		t.Errorf("port default = %q", cfg.Port)
	}
	if cfg.ListenAddr != "127.0.0.1" {
		t.Errorf("listen addr default = %q, want 127.0.0.1", cfg.ListenAddr)
	}
	if cfg.FrontendURL != "http://localhost:5173" {
		t.Errorf("frontend url default = %q", cfg.FrontendURL)
	}
}

func TestLoadFromListenAddrOverride(t *testing.T) {
	cfg, err := LoadFrom(env(map[string]string{"JWT_SECRET": validSecret, "LISTEN_ADDR": "0.0.0.0", "PORT": "9000"}))
	if err != nil {
		t.Fatal(err)
	}
	if cfg.ListenAddr != "0.0.0.0" || cfg.Port != "9000" {
		t.Fatalf("got %q:%q", cfg.ListenAddr, cfg.Port)
	}
}

func TestSecureCookies(t *testing.T) {
	cfg := &Config{FrontendURL: "https://example.com"}
	if !cfg.SecureCookies() {
		t.Error("https frontend should use Secure cookies")
	}
	cfg = &Config{FrontendURL: "http://localhost:5173"}
	if cfg.SecureCookies() {
		t.Error("http frontend should not use Secure cookies")
	}
}
