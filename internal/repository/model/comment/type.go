package comment

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

type JsonEncoder func(post models.Comment) ([]byte, error)
type SliceJsonEncoder func(post []models.Comment) ([]byte, error)

func SameIds(first models.Comment, second models.Comment) bool {
	return first == models.Comment(nil) && second == models.Comment(nil) || uuid.Equal(first.CommentId(), second.CommentId())
}

const PostRequired = "comment_post_required"
const PostIdAndPostFound = "comment_post_id_and_post"
const PostInvalidType = "comment_post_invalid_type"
const ContentRequired = "comment_content_required"
const AuthorNameRequired = "comment_author_name_required"
