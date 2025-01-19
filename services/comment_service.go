package services

import (
	"comments_api/models"
	"comments_api/repositories"
	"errors"
)

type CommentService struct {
    repo *repositories.CommentRepository
}

func NewCommentService(repo *repositories.CommentRepository) *CommentService {
    return &CommentService{
        repo: repo,
    }
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
    if err := validateComment(comment); err != nil {
        return err
    }
    return s.repo.Create(comment)
}

func (s *CommentService) GetAllComments() ([]models.Comment, error) {
    return s.repo.FindAll()
}

func (s *CommentService) DeleteComment(id int) error {
    return s.repo.Delete(id)
}

func validateComment(comment *models.Comment) error {
    if comment.UserID == 0 {
        return errors.New("userID is required")
    }
    if comment.Title == "" {
        return errors.New("title is required")
    }
    if comment.Body == "" {
        return errors.New("body is required")
    }
    return nil
}
