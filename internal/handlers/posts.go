package handlers

import (
	"net/http"

	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model/post"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/util"
)

type GetPosts struct {
	postRepository repository.Posts
	postsToJson post.SliceJsonEncoder
	logger *logger.Logger
}
func NewGetPosts(postRepository repository.Posts, postsToJson post.SliceJsonEncoder, logger *logger.Logger) *GetPosts {
	ctx := new(GetPosts)
	ctx.postRepository = postRepository
	ctx.postsToJson = postsToJson
	ctx.logger = logger
	return ctx
}

func (ctx *GetPosts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	posts, err := ctx.postRepository.GetAll()
	if err != nil {

	}
	bytes, err := ctx.postsToJson(posts)
	if err != nil {

	}
	util.WriteSafely(w, ctx.logger, bytes)
}
