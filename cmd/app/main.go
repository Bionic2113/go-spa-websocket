package main

import (
	"database/sql"
	"fmt"
	"sushkof/internal/database"
	"sushkof/internal/services"
	"sushkof/internal/transport"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
  var user, password, dbname string
  user = "root"
  password = "my_password"
  dbname = "sushkof"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		panic(err)
	}
  
  services.InitRepositories(database.UserRepoDB{DB: db}, database.TaskRepoDB{DB: db})
  fmt.Printf("users: %+v\n", database.UserRepoDB{DB: db}.FindAllActiveUsers())

  transport.StartServer(8080)
}
