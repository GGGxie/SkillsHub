package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"skillshub/internal/model"

	"github.com/gin-gonic/gin"
)

const (
	defaultPageSize = 12
	maxPageSize     = 50
	maxPage         = 100000
)

type SkillHandler struct {
	db    *sql.DB
	dedup *actionDedup
}

func NewSkillHandler(db *sql.DB, dedup *actionDedup) *SkillHandler {
	return &SkillHandler{db: db, dedup: dedup}
}

func (h *SkillHandler) ListSkills(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "12"))
	sort := c.DefaultQuery("sort", "hottest")
	category := c.Query("category")
	search := c.Query("search")

	if page < 1 {
		page = 1
	}
	if page > maxPage {
		page = maxPage
	}
	if pageSize < 1 {
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	offset := (page - 1) * pageSize

	query := "SELECT id, title, title_zh, description, desc_zh, category, icon, image, author_id, author_name, author_avatar, tags, skill_type, featured, likes, views, created_at, updated_at FROM skills WHERE 1=1"
	countQuery := "SELECT COUNT(*) FROM skills WHERE 1=1"
	args := []interface{}{}

	if category != "" {
		query += " AND category = ?"
		countQuery += " AND category = ?"
		args = append(args, category)
	}

	if search != "" {
		query += " AND (title LIKE ? OR description LIKE ? OR tags LIKE ?)"
		countQuery += " AND (title LIKE ? OR description LIKE ? OR tags LIKE ?)"
		s := "%" + search + "%"
		args = append(args, s, s, s)
	}

	var total int
	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	if err := h.db.QueryRow(countQuery, countArgs...).Scan(&total); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch skills"})
		return
	}

	switch sort {
	case "latest":
		query += " ORDER BY created_at DESC"
	default:
		query += " ORDER BY likes DESC, views DESC"
	}

	query += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	rows, err := h.db.Query(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch skills"})
		return
	}
	defer rows.Close()

	skills := []model.Skill{}
	for rows.Next() {
		var s model.Skill
		rows.Scan(&s.ID, &s.Title, &s.TitleZH, &s.Description, &s.DescZH, &s.Category, &s.Icon, &s.Image,
			&s.AuthorID, &s.AuthorName, &s.AuthorAvatar, &s.Tags, &s.SkillType, &s.Featured, &s.Likes, &s.Views,
			&s.CreatedAt, &s.UpdatedAt)
		skills = append(skills, s)
	}

	c.JSON(http.StatusOK, model.PaginatedResponse{
		Data:       skills,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	})
}

func (h *SkillHandler) GetFeaturedSkills(c *gin.Context) {
	rows, err := h.db.Query(`SELECT id, title, title_zh, description, desc_zh, category, icon, image, author_id, author_name, author_avatar, tags, skill_type, featured, likes, views, created_at, updated_at
		FROM skills WHERE featured = 1 ORDER BY likes DESC LIMIT 6`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch featured skills"})
		return
	}
	defer rows.Close()

	skills := []model.Skill{}
	for rows.Next() {
		var s model.Skill
		rows.Scan(&s.ID, &s.Title, &s.TitleZH, &s.Description, &s.DescZH, &s.Category, &s.Icon, &s.Image,
			&s.AuthorID, &s.AuthorName, &s.AuthorAvatar, &s.Tags, &s.SkillType, &s.Featured, &s.Likes, &s.Views,
			&s.CreatedAt, &s.UpdatedAt)
		skills = append(skills, s)
	}

	c.JSON(http.StatusOK, skills)
}

func (h *SkillHandler) GetSkill(c *gin.Context) {
	id := c.Param("id")

	key := fmt.Sprintf("view:skill:%s:%s", c.ClientIP(), id)
	if h.dedup.Allow(key, 1*time.Hour) {
		h.db.Exec("UPDATE skills SET views = views + 1 WHERE id = ?", id)
	}

	var s model.Skill
	err := h.db.QueryRow(`SELECT id, title, title_zh, description, desc_zh, category, icon, image, author_id, author_name, author_avatar, content, content_zh, tags, skill_type, featured, likes, views, created_at, updated_at
		FROM skills WHERE id = ?`, id).
		Scan(&s.ID, &s.Title, &s.TitleZH, &s.Description, &s.DescZH, &s.Category, &s.Icon, &s.Image,
			&s.AuthorID, &s.AuthorName, &s.AuthorAvatar, &s.Content, &s.ContentZH, &s.Tags, &s.SkillType, &s.Featured, &s.Likes, &s.Views,
			&s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h *SkillHandler) CreateSkill(c *gin.Context) {
	var req model.CreateSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("user_id")
	userName, _ := c.Get("name")

	result, err := h.db.Exec(`INSERT INTO skills (title, title_zh, description, desc_zh, category, icon, image, author_id, author_name, content, content_zh, tags, skill_type)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		req.Title, req.TitleZH, req.Description, req.DescZH, req.Category, req.Icon, req.Image,
		userID, userName, req.Content, req.ContentZH, req.Tags, req.SkillType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create skill"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": id, "message": "Skill created successfully"})
}

// LikeSkill toggles the current user's like on a skill. The existence check,
// the like row and the counter update run in one transaction so they either
// all apply or none do.
func (h *SkillHandler) LikeSkill(c *gin.Context) {
	skillID := c.Param("id")
	userID, _ := c.Get("user_id")

	fail := func(msg string, err error) {
		log.Printf("like skill %s: %s: %v", skillID, msg, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update like"})
	}

	tx, err := h.db.BeginTx(c.Request.Context(), nil)
	if err != nil {
		fail("begin", err)
		return
	}
	defer tx.Rollback() // no-op after a successful Commit

	var found int
	if err := tx.QueryRow("SELECT 1 FROM skills WHERE id = ?", skillID).Scan(&found); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
			return
		}
		fail("lookup skill", err)
		return
	}

	res, err := tx.Exec("DELETE FROM likes WHERE user_id = ? AND skill_id = ?", userID, skillID)
	if err != nil {
		fail("delete like", err)
		return
	}
	removed, err := res.RowsAffected()
	if err != nil {
		fail("rows affected", err)
		return
	}

	liked := removed == 0
	if liked {
		if _, err := tx.Exec("INSERT INTO likes (user_id, skill_id) VALUES (?, ?)", userID, skillID); err != nil {
			fail("insert like", err)
			return
		}
		_, err = tx.Exec("UPDATE skills SET likes = likes + 1 WHERE id = ?", skillID)
	} else {
		_, err = tx.Exec("UPDATE skills SET likes = MAX(likes - 1, 0) WHERE id = ?", skillID)
	}
	if err != nil {
		fail("update counter", err)
		return
	}
	if err := tx.Commit(); err != nil {
		fail("commit", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"liked": liked})
}

func (h *SkillHandler) GetCategories(c *gin.Context) {
	rows, err := h.db.Query("SELECT DISTINCT category FROM skills ORDER BY category")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	defer rows.Close()

	categories := []string{}
	for rows.Next() {
		var cat string
		rows.Scan(&cat)
		categories = append(categories, cat)
	}

	c.JSON(http.StatusOK, categories)
}

func (h *SkillHandler) GetComments(c *gin.Context) {
	skillID := c.Param("id")
	rows, err := h.db.Query("SELECT id, skill_id, user_id, user_name, avatar, content, created_at FROM comments WHERE skill_id = ? ORDER BY created_at DESC", skillID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch comments"})
		return
	}
	defer rows.Close()

	comments := []model.Comment{}
	for rows.Next() {
		var cm model.Comment
		rows.Scan(&cm.ID, &cm.SkillID, &cm.UserID, &cm.UserName, &cm.Avatar, &cm.Content, &cm.CreatedAt)
		comments = append(comments, cm)
	}

	c.JSON(http.StatusOK, comments)
}

func (h *SkillHandler) CreateComment(c *gin.Context) {
	skillID := c.Param("id")
	var req model.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var found int
	if err := h.db.QueryRow("SELECT 1 FROM skills WHERE id = ?", skillID).Scan(&found); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Skill not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	userID, _ := c.Get("user_id")
	userName, _ := c.Get("name")
	var avatar string
	h.db.QueryRow("SELECT avatar FROM users WHERE id = ?", userID).Scan(&avatar)

	result, err := h.db.Exec("INSERT INTO comments (skill_id, user_id, user_name, avatar, content) VALUES (?, ?, ?, ?, ?)",
		skillID, userID, userName, avatar, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	id, _ := result.LastInsertId()
	c.JSON(http.StatusCreated, gin.H{"id": id})
}
