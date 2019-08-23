package handlers

import (
	"encoding/json"
	"github.com/MNFGroup/openapimux"
	uuid "github.com/satori/go.uuid"
	models "github.com/shevchenkobn/blog-backend/internal/repository/model"
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
		panic(err)
	}
	bytes, err := ctx.postsToJson(posts)
	if err != nil {
		panic(err)
	}
	util.WriteSafely(w, ctx.logger, bytes)
}

type CreateOnePost struct {
	postRepository repository.Posts
	postToJson post.JsonEncoder
	logger *logger.Logger
}
func NewCreateOnePost(postRepository repository.Posts, postToJson post.JsonEncoder, logger *logger.Logger) *CreateOnePost {
	ctx := new(CreateOnePost)
	ctx.postRepository = postRepository
	ctx.postToJson = postToJson
	ctx.logger = logger
	return ctx
}
func (ctx *CreateOnePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var postSeed *models.PostSeed
	err := decoder.Decode(postSeed)
	if err != nil {
		panic(err)
	}

	p, err := ctx.postRepository.CreateOne(postSeed)
	if err != nil {
		panic(err)
	}
	bytes, err := ctx.postToJson(p)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	util.WriteSafely(w, ctx.logger, bytes)
}

type DeleteOnePost struct {
	postRepository repository.Posts
	logger *logger.Logger
}
func NewDeleteOnePost(postRepository repository.Posts, postToJson post.JsonEncoder, logger *logger.Logger) *DeleteOnePost {
	ctx := new(DeleteOnePost)
	ctx.postRepository = postRepository
	ctx.logger = logger
	return ctx
}
func (ctx *DeleteOnePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	postIdParam := openapimux.PathParam(r, "postId")
	postId, err := uuid.FromString(postIdParam)
	if err != nil {
		panic(err)
	}
	_, err = ctx.postRepository.DeleteOne(postId, false)
	if err != nil {
		panic(err)
	}
	util.WriteSafely(w, ctx.logger, []byte("{}"))
}
