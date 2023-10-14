package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/ricassiocosta/simplebank/util"
)

// GetDB returns a valid database connection
func GetDB() (*sql.DB, error) {
	var err error
	var dbConfig *util.DBConfig

	dbConfig, err = util.GetDBConfig()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open(dbConfig.DBDriver, dbConfig.DBSource)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
