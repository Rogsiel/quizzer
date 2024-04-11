package server

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setRouter() {
    router := gin.Default()
    router.POST("/signup", server.createAccount)
    router.POST("login", server.UserLogin)
    
    authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

    authRoutes.GET("/user/:user_id", server.getUserQuizList)
    authRoutes.POST("/quizzes", server.createQuiz)
    authRoutes.GET("/quizzes/:id", server.getQuizByID)
    authRoutes.POST("/quizzes/:id/answer", server.SendAnswer)
    
    server.router = router
}
