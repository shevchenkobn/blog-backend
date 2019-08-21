package db

import (
	"fmt"
	"github.com/go-pg/pg"

	"github.com/shevchenkobn/blog-backend/internal/services/config"
	"github.com/shevchenkobn/blog-backend/internal/services/logger"
	"github.com/shevchenkobn/blog-backend/internal/types"
)

type PostgreDB struct {
	db *pg.DB
	logger *logger.Logger
	onExit types.ExitHandler
	callback types.ExitHandlerCallback
}
func (p *PostgreDB) Db() *pg.DB {
	return p.db
}

func (p *PostgreDB) Close() {
	err := p.db.Close()
	if err != nil {
		p.logger.Errorf("Error closing DB: %s", err)
	}
	p.onExit.RemoveCallback(p.callback)
}

func NewPostgreDB(config config.Config, onExit types.ExitHandler, l *logger.Logger) *PostgreDB {
	connectConfig := &pg.Options{
		Addr: fmt.Sprintf("%s:%d", config.Db().Host(), config.Db().Port()),
		Database: config.Db().Database(),
		User: config.Db().User(),
	}
	if user := config.Db().User(); user != "" {
		connectConfig.User = user
	}
	if password := config.Db().Password(); password != "" {
		connectConfig.Password = password
	}
	p := &PostgreDB{
		db: pg.Connect(connectConfig),
		onExit: onExit,
		logger: l,
	}
	callback := func() {
		p.Close()
	}
	p.onExit.AddCallback(callback)
	p.callback = callback
	return p
}
