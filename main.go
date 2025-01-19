package main

import (
	"comments_api/config"
	"comments_api/repositories"
	"comments_api/routes"
	"comments_api/services"
	"log"
	"net/http"
)

func main() {
    db, err := config.InitializeDatabase()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }

    // Initialize dependencies
    commentRepo := repositories.NewCommentRepository(db)
    commentService := services.NewCommentService(commentRepo)
    
    // Initialize routes
    router := routes.InitializeRoutes(commentService)

    log.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
