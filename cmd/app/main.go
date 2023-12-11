package main

import (
	"database/sql"
	"fmt"
  "flag"
	"sushkof/internal/database"
	"sushkof/internal/services"
	"sushkof/internal/transport"
  
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  var user, password, dbname, host string
  // на маке к докер контейнеру другой докерконтейнер не хочет подключаться,
  // поэтому приходится указывать свой ip
  flag.StringVar(&host, "h", "localhost", "Need for macbook")
  flag.Parse()
  user = "root"
  password = "my_password"
  dbname = "sushkof"
  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", user, password, host, dbname))
	if err != nil {
		panic(err)
	}
  
  services.InitRepositories(database.UserRepoDB{DB: db}, database.TaskRepoDB{DB: db})
  fmt.Printf("users: %+v\n", database.UserRepoDB{DB: db}.FindAllActiveUsers())

  transport.StartServer(9080)
}
