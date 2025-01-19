package repositories

import (
	"comments_api/models"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type CommentRepository struct {
    db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
    return &CommentRepository{
        db: db,
    }
}

func (r *CommentRepository) Create(comment *models.Comment) error {
    result := r.db.Create(comment)
    if result.Error != nil {
        if strings.Contains(result.Error.Error(), "UNIQUE constraint failed") {
            return errors.New("comment already exists")
        }
        return result.Error
    }
    return nil
}

func (r *CommentRepository) FindAll() ([]models.Comment, error) {
    var comments []models.Comment
    err := r.db.Find(&comments).Error
    return comments, err
}

func (r *CommentRepository) Delete(id int) error {
    result := r.db.Delete(&models.Comment{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("comment not found")
    }
    return nil
}
