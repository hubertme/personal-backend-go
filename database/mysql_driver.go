package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMySQL() {
	fmt.Println("Postgres is running")

	dbTmp, err := sql.Open("mysql", os.Getenv("MYSQL_CONNECTION"))

	if err != nil {
		log.Fatal("MySQL failed to run:", err.Error())
	}

	DB = dbTmp
}
