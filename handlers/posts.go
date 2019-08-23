package handlers

import (
	"github.com/shevchenkobn/blog-backend/internal/util"
	"net/http"
)

func (ctx *HttpHandlers) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := ctx.postRepository.GetAll()
	if err != nil {

	}
	bytes, err := ctx.postsToJson(posts)
	if err != nil {

	}
	util.WriteSafely(w, ctx.logger, bytes)
}
