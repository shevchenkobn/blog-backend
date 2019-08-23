package di

import (
	"net/http"

	"github.com/shevchenkobn/blog-backend/internal/handlers"
	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/repository/model/comment"
	"github.com/shevchenkobn/blog-backend/internal/repository/model/post"
	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/services/db/pg"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/services/onexit"
	"github.com/shevchenkobn/blog-backend/internal/services/openapi"
	"github.com/shevchenkobn/blog-backend/internal/types"
)

var server *openapi.Server

func GetServer() *openapi.Server {
	if server == nil {
		server = openapi.NewServer(GetConfig(), GetHttpHandlers(), GetExitHandler(), GetLogger())
	}
	return server
}

func GetHttpHandlers() map[string]http.Handler {
	return map[string]http.Handler{
		"GetPosts":      handlers.NewGetPosts(GetPostRepository(), GetPostSliceJsonEncoder(), GetLogger()),
		"CreateOnePost": handlers.NewCreateOnePost(GetPostRepository(), GetPostJsonEncoder(), GetLogger()),
		"DeleteOnePost": handlers.NewDeleteOnePost(GetPostRepository(), GetLogger()),

		"GetCommentsForPost": handlers.NewGetCommentsForPost(GetCommentRepository(), GetPostRepository(), GetCommentSliceJsonEncoder(), GetLogger()),
		"CreateOneComment":   handlers.NewCreateOneComment(GetCommentRepository(), GetCommentJsonEncoder(), GetLogger()),
		"DeleteOneComment":   handlers.NewDeleteOneComment(GetCommentRepository(), GetLogger()),
	}
}

func GetPostJsonEncoder() post.JsonEncoder {
	return pg.PostToJson
}
func GetPostSliceJsonEncoder() post.SliceJsonEncoder {
	return pg.PostsToJson
}

func GetCommentJsonEncoder() comment.JsonEncoder {
	return pg.CommentToJson
}
func GetCommentSliceJsonEncoder() comment.SliceJsonEncoder {
	return pg.CommentsToJson
}

var postRepository repository.Posts

func GetPostRepository() repository.Posts {
	if postRepository == nil {
		postRepository = pg.NewPostRepository(GetPostgreDB())
	}
	return postRepository
}

var commentRepository repository.Comments

func GetCommentRepository() repository.Comments {
	if commentRepository == nil {
		commentRepository = pg.NewCommentRepository(GetPostgreDB())
	}
	return commentRepository
}

var postgreDb *db.PostgreDB

func GetPostgreDB() *db.PostgreDB {
	if postgreDb == nil {
		postgreDb = db.NewPostgreDB(GetConfig(), GetExitHandler(), GetLogger())
	}
	return postgreDb
}

var cachedConfig config.Config

func GetConfig() config.Config {
	if cachedConfig == nil {
		cachedConfig = config.GetConfig()
	}
	return cachedConfig
}

var exitHandler types.ExitHandler

func GetExitHandler() types.ExitHandler {
	if exitHandler == nil {
		exitHandler = onexit.NewExitHandler(GetLogger())
	}
	return exitHandler
}

var l *logger.Logger

func GetLogger() *logger.Logger {
	if l == nil {
		l = logger.NewLogger()
	}
	return l
}
