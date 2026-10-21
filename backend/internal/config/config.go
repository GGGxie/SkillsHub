package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	// MinJWTSecretLen is the minimum accepted length of JWT_SECRET.
	MinJWTSecretLen = 32
	// legacyJWTSecret was the hard-coded fallback in earlier versions; it must never be accepted.
	legacyJWTSecret = "your-secret-key-change-in-production"
)

type Config struct {
	Port               string
	ListenAddr         string
	DatabasePath       string
	JWTSecret          string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectURL  string
	FrontendURL        string
}

// Load reads configuration from the process environment.
func Load() (*Config, error) {
	return LoadFrom(os.Getenv)
}

// LoadFrom reads configuration through the supplied getenv function,
// which makes it usable in tests without touching the real environment.
func LoadFrom(getenv func(string) string) (*Config, error) {
	secret := getenv("JWT_SECRET")
	// Validate the trimmed value so a whitespace-only secret counts as missing,
	// but keep the raw value for signing so existing tokens stay valid.
	if err := validateJWTSecret(strings.TrimSpace(secret)); err != nil {
		return nil, err
	}
	get := func(key, fallback string) string {
		if v := getenv(key); v != "" {
			return v
		}
		return fallback
	}
	return &Config{
		Port:               get("PORT", "8080"),
		ListenAddr:         get("LISTEN_ADDR", "127.0.0.1"),
		DatabasePath:       get("DATABASE_PATH", "./skillshub.db"),
		JWTSecret:          secret,
		GoogleClientID:     get("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: get("GOOGLE_CLIENT_SECRET", ""),
		GoogleRedirectURL:  get("GOOGLE_REDIRECT_URL", "http://localhost:8080/api/auth/google/callback"),
		FrontendURL:        get("FRONTEND_URL", "http://localhost:5173"),
	}, nil
}

func validateJWTSecret(secret string) error {
	switch {
	case secret == "":
		return errors.New("JWT_SECRET is required but not set")
	case secret == legacyJWTSecret:
		return errors.New("JWT_SECRET must not use the legacy default value")
	case len(secret) < MinJWTSecretLen:
		return fmt.Errorf("JWT_SECRET is too short: need at least %d characters", MinJWTSecretLen)
	}
	return nil
}

// SecureCookies reports whether cookies should carry the Secure attribute.
func (c *Config) SecureCookies() bool {
	return strings.HasPrefix(c.FrontendURL, "https://")
}
