package main

import (
	"agnos-hospital/config"
	"agnos-hospital/internal/handler"
	"agnos-hospital/internal/middleware"
	"agnos-hospital/internal/repository"
	"agnos-hospital/internal/service"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.LoadConfig()

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Database not reachable:", err)
	}

	r := gin.Default()

	// Repositories
	staffRepo := repository.NewStaffRepository(db)
	patientRepo := repository.NewPatientRepository(db)

	// Services
	authService := service.NewAuthService(staffRepo, cfg.JWTSecret)
	patientService := service.NewPatientService(patientRepo)

	// Handlers
	staffHandler := handler.NewStaffHandler(authService)
	patientHandler := handler.NewPatientHandler(patientService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/staff/create", staffHandler.Register)
	r.POST("/staff/login", staffHandler.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/patient/search", patientHandler.Search)

	log.Println("Server running on :8080")
	r.Run(":8080")
}