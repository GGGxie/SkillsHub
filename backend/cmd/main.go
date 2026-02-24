package main

import (
	"log"

	"skillshub/internal/config"
	"skillshub/internal/handler"
	"skillshub/internal/middleware"
	"skillshub/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	db := service.InitDB(cfg.DatabasePath)
	defer db.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.FrontendURL, "http://localhost:5173", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	authHandler := handler.NewAuthHandler(db, cfg)
	skillHandler := handler.NewSkillHandler(db)
	articleHandler := handler.NewArticleHandler(db)

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

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
