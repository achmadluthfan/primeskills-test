package routes

import (
	"comments_api/handlers"
	"comments_api/services"
	"net/http"
)

func InitializeRoutes(commentService *services.CommentService) *http.ServeMux {
    router := http.NewServeMux()
    
    commentHandler := handlers.NewCommentHandler(commentService)

    router.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            commentHandler.CreateComment(w, r)
        case http.MethodGet:
            commentHandler.GetComments(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    router.HandleFunc("/comments/delete", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodDelete {
            commentHandler.DeleteComment(w, r)
        } else {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    return router
}
