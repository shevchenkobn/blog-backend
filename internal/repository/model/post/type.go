package post

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shevchenkobn/blog-backend/internal/repository/model"
)

//type Constructor interface {
//	NewPost(seed *PostSeed) PostField
//}
func SameIds(first models.Post, second models.Post) bool {
	return first == models.Post(nil) && second == models.Post(nil) || uuid.Equal(first.PostId(), second.PostId())
}

const ContentRequired = "post_content_required"
const AuthorNameRequired = "post_author_name_required"
