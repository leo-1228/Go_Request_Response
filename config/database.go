package config

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	// db, err := sql.Open("mysql", "root:''@tcp(localhost:3306)/go_products_master")
	// db, err := sql.Open("mysql", "root:''@/go_products_master?parseTime=true")
	// // db, err := sql.Open("mysql", "root:''@tcp(localhost:80)/go")
	cfg := mysql.Config{
		User:   os.Getenv("root"),
		Passwd: os.Getenv("root"),
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "go",
	}
	// Get a database handle.
	// var err error
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	log.Println("Database connected")
}
