package repository

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

type ModelError struct {
	Code string
}

func (e ModelError) Error() string {
	return "ModelError: " + e.Code
}

type Posts interface {
	GetAll() ([]models.Post, error)
	GetOne(postId uuid.UUID) (models.Post, error)
	CreateOne(post *models.PostSeed) (models.Post, error)
	DeleteOne(postId uuid.UUID, returning bool) (models.Post, error)
}

type Comments interface {
	GetAllForPost(postId uuid.UUID) ([]models.Comment, error)
	GetOne(commentId uuid.UUID) (models.Comment, error)
	CreateOne(comment *models.CommentSeed) (models.Comment, error)
	DeleteOne(commentId uuid.UUID, returning bool) (models.Comment, error)
}
