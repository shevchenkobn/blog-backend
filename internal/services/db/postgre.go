package db

import "github.com/go-pg/pg"

func GetPostgreConnection() *pg.DB {
	return pg.Connect(&pg.Options{

	})
}
