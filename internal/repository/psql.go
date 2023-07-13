package repository

import "github.com/go-pg/pg/v10"

var PG *pg.DB

func InitPostgresDB() {
	// 1. 链接
	PG = pg.Connect(&pg.Options{
		Addr:     "192.168.0.158:5432",
		User:     "root",
		Password: "msdnmm",
		Database: "db",
	})
}
