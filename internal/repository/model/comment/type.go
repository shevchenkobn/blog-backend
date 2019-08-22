package comment

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

//type Constructor interface {
//	NewComment(seed *CommentSeed) Comment
//}
func SameIds(first models.Comment, second models.Comment) bool {
	return first == models.Comment(nil) && second == models.Comment(nil) || uuid.Equal(first.CommentId(), second.CommentId())
}

const PostRequired = "comment_post_required"
const PostInvalidType = "comment_post_invalid_type"
const ContentRequired = "comment_content_required"
const AuthorNameRequired = "comment_author_name_required"
