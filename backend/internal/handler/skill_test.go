package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"skillshub/internal/model"
)

func validSkill() map[string]interface{} {
	return map[string]interface{}{
		"title":       "My skill",
		"description": "does things",
		"category":    "Coding",
		"content":     "# hello",
	}
}

func TestCreateSkillRequiresAuth(t *testing.T) {
	r, _, _ := testRouter(t)
	w := doJSON(r, http.MethodPost, "/api/skills", "", validSkill())
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status = %d, want 401", w.Code)
	}
}

func TestCreateSkillAcceptsValidAndBoundaryInput(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)

	if w := doJSON(r, http.MethodPost, "/api/skills", tok, validSkill()); w.Code != http.StatusCreated {
		t.Fatalf("valid: %d %s", w.Code, w.Body.String())
	}

	// exactly at the limits must still be accepted
	body := validSkill()
	body["title"] = strings.Repeat("t", 200)
	body["description"] = strings.Repeat("d", 2000)
	body["content"] = strings.Repeat("c", 50000)
	if w := doJSON(r, http.MethodPost, "/api/skills", tok, body); w.Code != http.StatusCreated {
		t.Fatalf("boundary: %d %s", w.Code, w.Body.String())
	}
	if n := count(t, db, "SELECT COUNT(*) FROM skills WHERE author_id = 1 AND title = ?", "My skill"); n != 1 {
		t.Fatalf("created skill rows = %d", n)
	}
}

func TestCreateSkillRejectsOversizedFields(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	before := count(t, db, "SELECT COUNT(*) FROM skills")

	limits := map[string]int{
		"title":       200,
		"title_zh":    200,
		"description": 2000,
		"desc_zh":     2000,
		"category":    100,
		"icon":        50,
		"image":       500,
		"content":     50000,
		"content_zh":  50000,
		"tags":        500,
		"skill_type":  50,
	}
	for field, max := range limits {
		t.Run(field, func(t *testing.T) {
			body := validSkill()
			body[field] = strings.Repeat("x", max+1)
			w := doJSON(r, http.MethodPost, "/api/skills", tok, body)
			if w.Code != http.StatusBadRequest {
				t.Fatalf("%s=%d chars: status %d, want 400", field, max+1, w.Code)
			}
		})
	}
	if after := count(t, db, "SELECT COUNT(*) FROM skills"); after != before {
		t.Fatalf("rejected requests inserted rows: %d -> %d", before, after)
	}
}

func TestOversizedBodyRejected(t *testing.T) {
	r, _, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	body := validSkill()
	body["content"] = strings.Repeat("a", 300*1024) // > 256KB
	w := doJSON(r, http.MethodPost, "/api/skills", tok, body)
	if w.Code != http.StatusRequestEntityTooLarge && w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want 413/400", w.Code)
	}
}

func TestCreateCommentLimitsAndTarget(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	path := "/api/skills/1/comments"

	if w := doJSON(r, http.MethodPost, path, "", map[string]string{"content": "hi"}); w.Code != http.StatusUnauthorized {
		t.Fatalf("anonymous comment: %d", w.Code)
	}
	if w := doJSON(r, http.MethodPost, path, tok, map[string]string{"content": "hi"}); w.Code != http.StatusCreated {
		t.Fatalf("valid comment: %d %s", w.Code, w.Body.String())
	}
	if w := doJSON(r, http.MethodPost, path, tok, map[string]string{"content": strings.Repeat("c", 2000)}); w.Code != http.StatusCreated {
		t.Fatalf("boundary comment: %d", w.Code)
	}
	if w := doJSON(r, http.MethodPost, path, tok, map[string]string{"content": strings.Repeat("c", 2001)}); w.Code != http.StatusBadRequest {
		t.Fatalf("oversize comment: %d, want 400", w.Code)
	}
	if w := doJSON(r, http.MethodPost, path, tok, map[string]string{"content": ""}); w.Code != http.StatusBadRequest {
		t.Fatalf("empty comment: %d, want 400", w.Code)
	}
	before := count(t, db, "SELECT COUNT(*) FROM comments")
	if w := doJSON(r, http.MethodPost, "/api/skills/999999/comments", tok, map[string]string{"content": "hi"}); w.Code != http.StatusNotFound {
		t.Fatalf("comment on missing skill: %d, want 404", w.Code)
	}
	if after := count(t, db, "SELECT COUNT(*) FROM comments"); after != before {
		t.Fatal("comment on missing skill was stored")
	}
}

func listSkills(t *testing.T, r http.Handler, query string) (int, model.PaginatedResponse) {
	t.Helper()
	w := doJSON(r, http.MethodGet, "/api/skills"+query, "", nil)
	var resp model.PaginatedResponse
	if w.Code == http.StatusOK {
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("decode: %v", err)
		}
	}
	return w.Code, resp
}

func TestListSkillsSearchIsParameterised(t *testing.T) {
	r, db, _ := testRouter(t)
	skills := count(t, db, "SELECT COUNT(*) FROM skills")
	users := count(t, db, "SELECT COUNT(*) FROM users")
	if skills == 0 {
		t.Fatal("seed produced no skills")
	}

	payloads := []string{
		"'; DROP TABLE skills; --",
		"' OR '1'='1",
		"%' UNION SELECT 1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18 --",
		`"; DELETE FROM users; --`,
	}
	for _, p := range payloads {
		for _, param := range []string{"search", "category"} {
			t.Run(param+"/"+p, func(t *testing.T) {
				code, resp := listSkills(t, r, "?"+param+"="+url.QueryEscape(p))
				if code != http.StatusOK {
					t.Fatalf("status = %d, want 200", code)
				}
				if resp.Total != 0 {
					t.Fatalf("injection-style input matched %d rows", resp.Total)
				}
			})
		}
	}
	if got := count(t, db, "SELECT COUNT(*) FROM skills"); got != skills {
		t.Fatalf("skills changed: %d -> %d", skills, got)
	}
	if got := count(t, db, "SELECT COUNT(*) FROM users"); got != users {
		t.Fatalf("users changed: %d -> %d", users, got)
	}
	// and the table still answers normal queries
	if code, resp := listSkills(t, r, ""); code != http.StatusOK || resp.Total != skills {
		t.Fatalf("normal list after payloads: code=%d total=%d want %d", code, resp.Total, skills)
	}
}

func TestListSkillsPageSizeClamped(t *testing.T) {
	r, _, _ := testRouter(t)
	cases := map[string]int{
		"":                  12,
		"?page_size=5":      5,
		"?page_size=50":     50,
		"?page_size=51":     50,
		"?page_size=100000": 50,
		"?page_size=0":      12,
		"?page_size=-3":     12,
		"?page_size=abc":    12,
	}
	for q, want := range cases {
		t.Run(q, func(t *testing.T) {
			code, resp := listSkills(t, r, q)
			if code != http.StatusOK {
				t.Fatalf("status %d", code)
			}
			if resp.PageSize != want {
				t.Fatalf("page_size = %d, want %d", resp.PageSize, want)
			}
			if len(resp.Data.([]interface{})) > want {
				t.Fatalf("returned more than %d rows", want)
			}
		})
	}
}

func TestListSkillsPageBounds(t *testing.T) {
	r, _, _ := testRouter(t)
	for _, q := range []string{"?page=0", "?page=-5", "?page=abc", "?page=9223372036854775807", "?page=99999999999999999999"} {
		t.Run(q, func(t *testing.T) {
			code, resp := listSkills(t, r, q)
			if code != http.StatusOK {
				t.Fatalf("status %d", code)
			}
			if resp.Page < 1 || resp.Page > maxPage {
				t.Fatalf("page = %d out of [1,%d]", resp.Page, maxPage)
			}
		})
	}
}

func skillLikes(t *testing.T, r http.Handler, id int) int {
	t.Helper()
	w := doJSON(r, http.MethodGet, "/api/skills/"+strconv.Itoa(id), "", nil)
	var s model.Skill
	if err := json.Unmarshal(w.Body.Bytes(), &s); err != nil || w.Code != http.StatusOK {
		t.Fatalf("get skill %d: %d %s", id, w.Code, w.Body.String())
	}
	return s.Likes
}

func TestLikeSkillToggles(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	base := skillLikes(t, r, 1)

	if w := doJSON(r, http.MethodPost, "/api/skills/1/like", "", nil); w.Code != http.StatusUnauthorized {
		t.Fatalf("anonymous like: %d", w.Code)
	}

	w := doJSON(r, http.MethodPost, "/api/skills/1/like", tok, nil)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"liked":true`) {
		t.Fatalf("first like: %d %s", w.Code, w.Body.String())
	}
	if got := skillLikes(t, r, 1); got != base+1 {
		t.Fatalf("likes after like = %d, want %d", got, base+1)
	}
	if n := count(t, db, "SELECT COUNT(*) FROM likes WHERE user_id = 1 AND skill_id = 1"); n != 1 {
		t.Fatalf("like rows = %d", n)
	}

	w = doJSON(r, http.MethodPost, "/api/skills/1/like", tok, nil)
	if w.Code != http.StatusOK || !strings.Contains(w.Body.String(), `"liked":false`) {
		t.Fatalf("unlike: %d %s", w.Code, w.Body.String())
	}
	if got := skillLikes(t, r, 1); got != base {
		t.Fatalf("likes after unlike = %d, want %d", got, base)
	}
	if n := count(t, db, "SELECT COUNT(*) FROM likes WHERE user_id = 1 AND skill_id = 1"); n != 0 {
		t.Fatalf("like rows after unlike = %d", n)
	}
}

func TestLikeSkillMissingTargetIs404(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	for _, id := range []string{"999999", "0", "-1", "abc"} {
		t.Run(id, func(t *testing.T) {
			w := doJSON(r, http.MethodPost, "/api/skills/"+id+"/like", tok, nil)
			if w.Code != http.StatusNotFound {
				t.Fatalf("status = %d, want 404 (%s)", w.Code, w.Body.String())
			}
		})
	}
	if n := count(t, db, "SELECT COUNT(*) FROM likes"); n != 0 {
		t.Fatalf("likes rows created for missing skills: %d", n)
	}
}

func TestLikeSkillIsAtomicOnFailure(t *testing.T) {
	r, db, cfg := testRouter(t)
	tok := tokenFor(t, cfg, 1)
	base := skillLikes(t, r, 1)

	// Make the counter update fail: the like row insert must be rolled back.
	if _, err := db.Exec(`CREATE TRIGGER fail_like_update BEFORE UPDATE OF likes ON skills
		BEGIN SELECT RAISE(ABORT, 'boom'); END`); err != nil {
		t.Fatal(err)
	}
	w := doJSON(r, http.MethodPost, "/api/skills/1/like", tok, nil)
	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d, want 500", w.Code)
	}
	if n := count(t, db, "SELECT COUNT(*) FROM likes"); n != 0 {
		t.Fatalf("like row survived failed transaction: %d", n)
	}
	if _, err := db.Exec("DROP TRIGGER fail_like_update"); err != nil {
		t.Fatal(err)
	}
	if got := skillLikes(t, r, 1); got != base {
		t.Fatalf("likes changed: %d -> %d", base, got)
	}
}
