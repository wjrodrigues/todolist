package test

import (
	"database/sql"
	"fmt"
	"testing"
	"todolist/configs"
	"todolist/pkg/file"

	_ "github.com/lib/pq"
)

func Conn(t *testing.T) *sql.DB {
	configs, _ := configs.LoadEnv(".", file.Path(".env"))

	dbName := fmt.Sprintf("%s_test", configs.DBName)
	db, _ := sql.Open(
		configs.DBDriver,
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			configs.DBHost, configs.DBPort, configs.DBUser, configs.DBPassword, dbName))

	return db
}
