package di

import (
	"github.com/shevchenkobn/blog-backend/internal/repository"
	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/services/db/pg"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/services/onexit"
	"github.com/shevchenkobn/blog-backend/internal/types"
)

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
