package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ahfrd/grpc/micro-topup/config"
)

// Database is a
type Database struct{}

// ConnectDB is a
func (o Database) ConnectDB() (*sql.DB, error) {
	c, err := config.LoadConfig()
	DB := c.DB
	URI := c.DBUrl
	fmt.Println(DB)
	fmt.Println(URI)
	// db, err := apmsql.Open(DB, URI)
	db, err := sql.Open(DB, URI)

	if err != nil {
		fmt.Println("lkk")
		fmt.Println(err)
		return nil, fmt.Errorf("failed connection to DB : %v", err)
	}

	return db, nil
}
