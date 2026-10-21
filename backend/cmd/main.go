package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"skillshub/internal/config"
	"skillshub/internal/handler"
	"skillshub/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	// Release mode unless the operator explicitly chose a mode via GIN_MODE.
	if os.Getenv(gin.EnvGinMode) == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	db := service.InitDB(cfg.DatabasePath)
	defer db.Close()

	r, err := handler.NewRouter(db, cfg)
	if err != nil {
		log.Fatalf("Failed to build router: %v", err)
	}

	addr := net.JoinHostPort(cfg.ListenAddr, cfg.Port)
	srv := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 10 * time.Second,
	}
	log.Printf("Server starting on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
