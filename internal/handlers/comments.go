package handlers

import (
	"encoding/json"
	"github.com/shevchenkobn/blog-backend/internal/repository"
	models "github.com/shevchenkobn/blog-backend/internal/repository/model"
	"github.com/shevchenkobn/blog-backend/internal/repository/model/comment"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
	"github.com/shevchenkobn/blog-backend/internal/util"
	"net/http"
)

type GetCommentsForPost struct {
	commentRepository repository.Comments
	postRepository    repository.Posts
	commentsToJson    comment.SliceJsonEncoder
	logger            *logger.Logger
}

func NewGetCommentsForPost(commentRepository repository.Comments, postRepository repository.Posts, commentsToJson comment.SliceJsonEncoder, logger *logger.Logger) *GetCommentsForPost {
	ctx := new(GetCommentsForPost)
	ctx.commentRepository = commentRepository
	ctx.postRepository = postRepository
	ctx.commentsToJson = commentsToJson
	ctx.logger = logger
	return ctx
}
func (ctx *GetCommentsForPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	postId := util.GetUuidPathParam(r, "postId")
	p, err := ctx.postRepository.GetOne(postId)
	if err != nil {
		panic(err)
	}
	if p == nil {
		panic(types.NewLogicError(types.ErrorPostNotFound))
	}
	comments, err := ctx.commentRepository.GetAllForPost(postId)
	if err != nil {
		panic(err)
	}
	bytes, err := ctx.commentsToJson(comments)
	if err != nil {
		panic(err)
	}
	util.WriteSafely(w, ctx.logger, bytes)
}

type CreateOneComment struct {
	commentRepository repository.Comments
	commentToJson     comment.JsonEncoder
	logger            *logger.Logger
}

func NewCreateOneComment(commentRepository repository.Comments, commentToJson comment.JsonEncoder, logger *logger.Logger) *CreateOneComment {
	ctx := new(CreateOneComment)
	ctx.commentRepository = commentRepository
	ctx.commentToJson = commentToJson
	ctx.logger = logger
	return ctx
}
func (ctx *CreateOneComment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var commentSeed = &models.CommentSeed{}
	err := decoder.Decode(commentSeed)
	if err != nil {
		panic(err)
	}
	commentSeed.PostId = util.GetUuidPathParam(r, "postId")

	c, err := ctx.commentRepository.CreateOne(commentSeed)
	if err != nil {
		panic(err)
	}
	bytes, err := ctx.commentToJson(c)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	util.WriteSafely(w, ctx.logger, bytes)
}

type DeleteOneComment struct {
	commentRepository repository.Comments
	logger            *logger.Logger
}

func NewDeleteOneComment(commentRepository repository.Comments, logger *logger.Logger) *DeleteOneComment {
	ctx := new(DeleteOneComment)
	ctx.commentRepository = commentRepository
	ctx.logger = logger
	return ctx
}
func (ctx *DeleteOneComment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := ctx.commentRepository.DeleteOne(util.GetUuidPathParam(r, "commentId"), false)
	if err != nil {
		panic(err)
	}
	util.WriteSafely(w, ctx.logger, []byte("{}"))
}
