package service

import (
	"fmt"
	"sync/atomic"
	"testing"
)

var memCounter int64

// memDSN returns a unique in-memory SQLite path (shared cache so every pooled
// connection sees the same database).
func memDSN() string {
	n := atomic.AddInt64(&memCounter, 1)
	return fmt.Sprintf("file:servicetest%d?mode=memory&cache=shared", n)
}

func TestOpenDBEnablesForeignKeysAndSeeds(t *testing.T) {
	db, err := OpenDB(memDSN())
	if err != nil {
		t.Fatalf("OpenDB: %v", err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	var fk int
	if err := db.QueryRow("PRAGMA foreign_keys").Scan(&fk); err != nil || fk != 1 {
		t.Fatalf("foreign_keys pragma = %d (err=%v), want 1", fk, err)
	}

	for table, min := range map[string]int{"users": 1, "skills": 1, "articles": 1} {
		var n int
		if err := db.QueryRow("SELECT COUNT(*) FROM " + table).Scan(&n); err != nil {
			t.Fatal(err)
		}
		if n < min {
			t.Errorf("table %s has %d rows after seed, want >= %d", table, n, min)
		}
	}

	// Every seeded skill must reference an existing author.
	var orphans int
	db.QueryRow(`SELECT COUNT(*) FROM skills WHERE author_id NOT IN (SELECT id FROM users)`).Scan(&orphans)
	if orphans != 0 {
		t.Errorf("%d orphan skills", orphans)
	}
	if rows, err := db.Query("PRAGMA foreign_key_check"); err == nil {
		defer rows.Close()
		if rows.Next() {
			t.Error("foreign_key_check reported violations after seed")
		}
	}
}

func TestForeignKeysAreEnforced(t *testing.T) {
	db, err := OpenDB(memDSN())
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.SetMaxOpenConns(1)

	_, err = db.Exec(`INSERT INTO comments (skill_id, user_id, content) VALUES (999999, 1, 'x')`)
	if err == nil {
		t.Fatal("expected FK violation for comment on a missing skill")
	}
}

func TestOpenDBIsIdempotent(t *testing.T) {
	dsn := memDSN()
	db, err := OpenDB(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var before int
	db.QueryRow("SELECT COUNT(*) FROM skills").Scan(&before)
	// Re-running schema + seed on an already populated DB must not duplicate data.
	if err := initSchemaAndSeed(db); err != nil {
		t.Fatal(err)
	}
	var after int
	db.QueryRow("SELECT COUNT(*) FROM skills").Scan(&after)
	if before != after {
		t.Fatalf("seed not idempotent: %d -> %d", before, after)
	}
}
