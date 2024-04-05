package main

import (
	"database/sql"
	"log"

	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/service/server"
	"github.com/rogsiel/quizzer/internal/util"
)

func main() {
    config, err := util.LoadConfig(".")
    if err != nil {
	log.Fatal("Can't load environment variables:", err)
    }
    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
	log.Fatal("Can't connect to Database:", err)
    }

    store := db.NewStore(conn)
    server := server.NewServer(store)

    err = server.Start(config.ServerAddress)
    if err != nil {
	log.Fatal("can't start server:", err)
    }
}
