package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/rogsiel/quizzer/internal/auth"
	db "github.com/rogsiel/quizzer/internal/database"
	"github.com/rogsiel/quizzer/internal/service/mail"
	"github.com/rogsiel/quizzer/internal/service/otp"
)

type createAccountReq struct {
    UserName	string  `json:"user_name" binding:"required,alphanum"`
    Email	string  `json:"email" binding:"required,email"`
    Password	string	`json:"password" binding:"required,min=8"`
}
type userResponse struct {
    UserName		string	    `json:"user_name"`
    Email		string	    `json:"email"`
    PasswrordChangedAt	time.Time   `json:"password_changed_at"`
    CreatedAt		time.Time   `json:"created_at"`
}

func newUserResponse(user db.User) userResponse{
    return userResponse{
	UserName: user.UserName,
	Email: user.Email,
	PasswrordChangedAt: user.PasswordChangedAt,
	CreatedAt: user.CreatedAt,
    }
}

func (server *Server) createAccount(ctx *gin.Context) {
    var req createAccountReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    
    hashedPassword, err := auth.HashPassword(req.Password)
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err)) 
	return
    }
    arg := db.CreateUserParams{
	UserName: req.UserName,
	Email: req.Email,
	HashedPassword: hashedPassword,	
    } 

    account, err := server.store.CreateUser(ctx, arg)
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    
    otp := server.otpManager.NewEmailVerificationOTP(account.Email)
    
    err = server.store.CreateOTPTx(ctx, db.CreateOTPTxParams{
	Email: otp.Email,
	OtpCode: otp.OtpCode,
	OtpType: otp.OtpType,
    })
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    sender := mail.NewEmail()
    sender.SendWelcomeEmail(mail.NewUserInfo{
	UserName: account.UserName,
	Email: account.Email,
	OtpCode: otp.OtpCode,
    })

    res := newUserResponse(account)
    ctx.JSON(http.StatusOK, res)
}

type userLoginReq struct {
    UserName	string  `json:"user_name" binding:"required,alphanum"`
    Password	string	`json:"password" binding:"required,min=8"`
}

type userLoginRes struct {
    SessionID		    uuid.UUID		`json:"session_id"`
    AccessToken		    string		`json:"access_token"`
    AccessTokenExpiresAt    time.Time		`json:"access_token_expires_at"`
    RefreshToken	    string		`json:"refresh_token"`
    RefreshTokenExpiresAt   time.Time		`json:"refresh_token_expires_at"`
    
    User		    userResponse	`json:"user"`
}

func (server *Server) UserLogin(ctx *gin.Context) {
    var req userLoginReq
    if err := ctx.ShouldBindJSON(&req); err != nil {
    	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    } 
    
    user, err := server.store.GetUser(ctx, req.UserName)
    if err != nil {
	if err == sql.ErrNoRows {
	    ctx.JSON(http.StatusNotFound, errorResponse(err))
	    return
	}
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }

    err = auth.CheckPassword(req.Password, user.HashedPassword)
    if err != nil {
	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	return
    }

    accessToken, accessPayload, err := server.tokenMaker.CreateToken(
	user.UserName,
	server.config.AccessTokenDuration,
    )
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
	user.UserName,
	server.config.RefreshTokenDuration,
    )
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    
    session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
	ID: refreshPayload.ID,
	UserName: user.UserName,
	RefreshToken: refreshToken,
	UserAgent: ctx.Request.UserAgent(),
	ClientIp: ctx.ClientIP(),
	IsBlocked: false,
	ExpiresAt: refreshPayload.ExpiredAt,
    })
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    

    res := userLoginRes{
	SessionID:		session.ID,
	AccessToken:		accessToken,
	AccessTokenExpiresAt:	accessPayload.ExpiredAt,
	RefreshToken:		refreshToken,
	RefreshTokenExpiresAt:	refreshPayload.ExpiredAt,
	User:			newUserResponse(user),
    }

    ctx.JSON(http.StatusOK, res)
}

type emailVerifyReq struct {
    Email   string  `uri:"email" binding:"required"`
    OtpCode string  `uri:"otp_code" binding:"required"`
    OtpType string  `uri:"otp_type" binding:"required"`
}

func (server *Server) emailVerify(ctx *gin.Context) {
    var req emailVerifyReq
    if err := ctx.ShouldBindUri(&req); err != nil {
    	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    record, err := server.store.GetOTP(ctx, db.GetOTPParams{
	Email: req.Email,
	OtpCode: req.OtpCode,
	OtpType: req.OtpType,
    })
    if err != nil {
	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	return
    }
    otp := otp.OTP{
	ID: record.ID,
	Email: record.Email,
	OtpCode: record.OtpCode,
	OtpType: record.OtpType,
	IsUsed: record.IsUsed,
	CreatedAt: record.CreatedAt,
	ExpiredAt: record.ExpiredAt,
    }
    
    err = otp.VerifyOTP()
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    
    username, err := server.store.VerifyEmailTx(ctx, db.VerifyEmailTxParams{
	OtpID: otp.ID,
	Email: otp.Email,
	OtpCode: otp.OtpCode,
    })
    if err != nil {
	ctx.JSON(http.StatusBadRequest, errorResponse(err))
	return
    }
    
    ctx.JSON(http.StatusOK, username)
}
