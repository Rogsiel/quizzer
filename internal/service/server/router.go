package server

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) setRouter() {
    router := gin.Default()
    router.POST("/signup", server.createAccount)
    
    router.POST("/token/renew_access", server.renewAccessToken)
    
    router.GET("/verify-email", server.emailVerify)
     
    loginRoutes := router.Group("/login")
    loginRoutes.POST("/", server.UserLogin)
    loginRoutes.POST("/forgot-password",server.forgotPassword)
    loginRoutes.POST("/reset-password",server.resetPassword)

    authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
    authRoutes.GET("/user/:user_id", server.getUserQuizList)
    authRoutes.POST("/quizzes", server.createQuiz)
    authRoutes.GET("/quizzes/:id", server.getQuizByID)
    authRoutes.POST("/quizzes/:id/answer", server.SendAnswer)

    server.router = router
}
