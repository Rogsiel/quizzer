package main

import (
	"database/sql"
	"log"

	"github.com/rogsiel/quizzer/config"
	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/service/server"
)

func main() {
    config, err := config.LoadConfig(".")
    if err != nil {
	log.Fatal("Can't load environment variables:", err)
    }
    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
	log.Fatal("Can't connect to Database:", err)
    }

    store := db.NewStore(conn)
    server, err := server.NewServer(config, *store)
    if err != nil {
	log.Fatal("can't initiate server:", err)
    }

    err = server.Start(config.ServerAddress)
    if err != nil {
	log.Fatal("can't start server:", err)
    }
}
