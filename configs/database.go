package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func DBConnection(configs env) *sql.DB {
	stringConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, configs.DBName)

	db, err := sql.Open(configs.DBDriver, stringConn)

	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
