package db

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var dbInstance *bun.DB

func InitDB() *bun.DB {
	dsn := "postgres://postgres:postgres@localhost:5432/learn-go?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	dbInstance = bun.NewDB(sqldb, pgdialect.New())

	return dbInstance
}

// Getconn return database connection instance
func GetConn() *bun.DB {
	return dbInstance
}
