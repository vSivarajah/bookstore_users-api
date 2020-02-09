package users_db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("root:root@tcp(0.0.0.0:3306)/users_db")

	var err error
	Client, err := sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		fmt.Println("should panic")
		panic(err)
	}
	log.Println("database succesfully configured")

}
