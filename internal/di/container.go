package di

import (
	"os"
	"syscall"

	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/db"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/services/onexit"
	"github.com/shevchenkobn/blog-backend/internal/types"
)

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
		exitHandler = onexit.NewExitHandler(GetLogger(), GetExitSignals())
	}
	return exitHandler
}

var exitSignals = []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP}
func GetExitSignals() []os.Signal {
	return exitSignals
}

var l *logger.Logger
func GetLogger() *logger.Logger {
	if l == nil {
		l = logger.NewLogger()
	}
	return l
}
