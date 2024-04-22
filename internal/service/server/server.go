package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rogsiel/quizzer/config"
	"github.com/rogsiel/quizzer/internal/auth/token"
	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/service/otp"
)

type Server struct {
    store	db.Store
    tokenMaker	token.Maker	
    otpManager	otp.OTPManager
    router	*gin.Engine
    config	config.Config
}

func NewServer(config config.Config,store db.Store) (*Server, error) {
    tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
    if err != nil {
	return nil, fmt.Errorf("can't initiate token maker: %w", err)
    }
    
    otpManager := otp.NewOTPManager()

    server := &Server{
	store: store,
	tokenMaker: tokenMaker,
	otpManager: otpManager,
	config: config,
    }
    server.setRouter()  
    return server, nil
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}
