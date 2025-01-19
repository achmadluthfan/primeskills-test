package handlers

import (
	"comments_api/models"
	"comments_api/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type CommentHandler struct {
    service *services.CommentService
}

func NewCommentHandler(service *services.CommentService) *CommentHandler {
    return &CommentHandler{
        service: service,
    }
}

func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
    var comment models.Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    if err := h.service.CreateComment(&comment); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comment)
}

func (h *CommentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
    comments, err := h.service.GetAllComments()
    if err != nil {
        http.Error(w, "Failed to retrieve comments", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(comments)
}

func (h *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing comment ID", http.StatusBadRequest)
        return
    }

    intId, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    if err := h.service.DeleteComment(intId); err != nil {
        if err.Error() == "comment not found" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Comment deleted successfully"))
}
