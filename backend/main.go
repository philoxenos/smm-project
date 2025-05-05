package main

import (
	"log"
	"os"

	"backend/controllers"
	"backend/middleware"
	"backend/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or error loading .env")
	}

	// Get frontend URL from env
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173" // default for Vite
	}

	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Initialize MySQL connection
	models.InitDB()

	// TODO: Setup Microsoft OAuth (see controllers/auth)

	// Route groups
	authRoutes := r.Group("/auth")
	{
		authRoutes.GET("/login", controllers.MicrosoftLogin)
		authRoutes.GET("/callback", controllers.MicrosoftCallback)
		authRoutes.GET("/microsoft/callback", controllers.MicrosoftCallback) // Fix for Microsoft redirect
		authRoutes.GET("/logout", controllers.Logout)
	}

	postRoutes := r.Group("/posts")
	{
		postRoutes.GET("/", controllers.GetPosts)
		postRoutes.POST("/", middleware.JWTAuthMiddleware(), controllers.CreatePost)
		postRoutes.PUT(":id", middleware.JWTAuthMiddleware(), controllers.UpdatePost)
		postRoutes.DELETE(":id", middleware.JWTAuthMiddleware(), controllers.DeletePost)
		postRoutes.POST(":id/like", middleware.JWTAuthMiddleware(), controllers.LikePost)
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Backend running on port %s", port)
	r.Run(":" + port)
}
