package server

import (
	"github.com/gin-gonic/gin"
	db "github.com/rogsiel/quizzer/internal/database"
)

type Server struct {
    store *db.Store
    router *gin.Engine
}

func NewServer(store *db.Store) *Server {
    server := &Server{store: store}
    router := gin.Default()
    
    router.POST("/signup", server.createAccount)
    router.GET("/search/:id",server.getUserQuizList)
    router.POST("/quiz/new", server.createQuiz)
    router.GET("/quiz/:id", server.getQuizByID)
    server.router = router
    return server
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}
