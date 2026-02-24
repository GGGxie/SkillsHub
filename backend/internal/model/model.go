package model

import "time"

type User struct {
	ID        int64     `json:"id"`
	GoogleID  string    `json:"google_id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}

type Skill struct {
	ID           int64     `json:"id"`
	Title        string    `json:"title"`
	TitleZH      string    `json:"title_zh"`
	Description  string    `json:"description"`
	DescZH       string    `json:"desc_zh"`
	Category     string    `json:"category"`
	Icon         string    `json:"icon"`
	Image        string    `json:"image"`
	AuthorID     int64     `json:"author_id"`
	AuthorName   string    `json:"author_name"`
	AuthorAvatar string    `json:"author_avatar"`
	Content      string    `json:"content"`
	ContentZH    string    `json:"content_zh"`
	Tags         string    `json:"tags"`
	SkillType    string    `json:"skill_type"`
	Featured     bool      `json:"featured"`
	Likes        int       `json:"likes"`
	Views        int       `json:"views"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Article struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	TitleZH     string    `json:"title_zh"`
	Description string    `json:"description"`
	DescZH      string    `json:"desc_zh"`
	Category    string    `json:"category"`
	Content     string    `json:"content"`
	ContentZH   string    `json:"content_zh"`
	AuthorID    int64     `json:"author_id"`
	AuthorName  string    `json:"author_name"`
	Views       int       `json:"views"`
	CreatedAt   time.Time `json:"created_at"`
}

type Comment struct {
	ID        int64     `json:"id"`
	SkillID   int64     `json:"skill_id"`
	UserID    int64     `json:"user_id"`
	UserName  string    `json:"user_name"`
	Avatar    string    `json:"avatar"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateSkillRequest struct {
	Title       string `json:"title" binding:"required"`
	TitleZH     string `json:"title_zh"`
	Description string `json:"description" binding:"required"`
	DescZH      string `json:"desc_zh"`
	Category    string `json:"category" binding:"required"`
	Icon        string `json:"icon"`
	Image       string `json:"image"`
	Content     string `json:"content" binding:"required"`
	ContentZH   string `json:"content_zh"`
	Tags        string `json:"tags"`
	SkillType   string `json:"skill_type"`
}

type CreateCommentRequest struct {
	Content string `json:"content" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}
