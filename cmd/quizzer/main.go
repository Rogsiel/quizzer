package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/service/server"
)

func main() {
    err := godotenv.Load()
    if err != nil {
	log.Fatal("Can't load environment variables:", err)
    }
    conn, err := sql.Open(os.Getenv("dbDriver"), os.Getenv("dbSource"))
    if err != nil {
	log.Fatal("Can't connect to Database:", err)
    }

    store := db.NewStore(conn)
    server := server.NewServer(store)

    err = server.Start(os.Getenv("serverAddress"))
    if err != nil {
	log.Fatal("can't start server:", err)
    }
}
