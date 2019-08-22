package repository

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

type ModelError struct {
	Code string
}
func (e ModelError) Error() string {
	return "ModelError:" + e.Code
}

type Posts interface {
	GetAll() ([]models.Post, error)
	//CreateOne(post post.PostSeed) (post.Post, error)
	//DeleteOne(postId uuid.UUID, returning bool) (post.Post, error)
}

type Comments interface {
	GetAllForPost(postId uuid.UUID) ([]models.Comment, error)
	//CreateOne(comment comment.PostSeed) (comment.Comment, error)
	//DeleteOne(commentId uuid.UUID, returning bool) (comment.Comment, error)
}
