package handler

import (
	"database/sql"
	"fmt"

	"skillshub/internal/config"
	"skillshub/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// MaxRequestBodyBytes is the largest request body the API accepts.
const MaxRequestBodyBytes = 256 * 1024

// trustedProxies lists the only peers whose X-Forwarded-For header is honoured
// (the local reverse proxy).
var trustedProxies = []string{"127.0.0.1"}

// NewRouter builds the HTTP router with all middleware and routes attached.
func NewRouter(db *sql.DB, cfg *config.Config) (*gin.Engine, error) {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	if err := r.SetTrustedProxies(trustedProxies); err != nil {
		return nil, fmt.Errorf("set trusted proxies: %w", err)
	}
	r.Use(middleware.MaxBodyBytes(MaxRequestBodyBytes))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL, "http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	dedup := NewDedup()
	authHandler := NewAuthHandler(db, cfg)
	skillHandler := NewSkillHandler(db, dedup)
	articleHandler := NewArticleHandler(db, dedup)

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.GET("/google", authHandler.GoogleLogin)
			auth.GET("/google/callback", authHandler.GoogleCallback)
			auth.POST("/google/token", authHandler.GoogleTokenLogin)
			auth.GET("/me", middleware.AuthRequired(cfg.JWTSecret), authHandler.GetMe)
		}

		skills := api.Group("/skills")
		{
			skills.GET("", skillHandler.ListSkills)
			skills.GET("/featured", skillHandler.GetFeaturedSkills)
			skills.GET("/categories", skillHandler.GetCategories)
			skills.GET("/:id", skillHandler.GetSkill)
			skills.POST("", middleware.AuthRequired(cfg.JWTSecret), skillHandler.CreateSkill)
			skills.POST("/:id/like", middleware.AuthRequired(cfg.JWTSecret), skillHandler.LikeSkill)
			skills.GET("/:id/comments", skillHandler.GetComments)
			skills.POST("/:id/comments", middleware.AuthRequired(cfg.JWTSecret), skillHandler.CreateComment)
		}

		articles := api.Group("/articles")
		{
			articles.GET("", articleHandler.ListArticles)
			articles.GET("/:id", articleHandler.GetArticle)
		}
	}
	return r, nil
}
