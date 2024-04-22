package otp

import (
	"math/rand"
	"strconv"
	"time"
)

const (
    EmailVerificationOTP    string = "email_verification"
    PasswordChangeOTP       string = "password_change"
)

func generateOTPCode() string {
    return strconv.Itoa(rand.Intn(900000) + 100000)
}

type OTP struct {
    ID        int64     `json:"id"`
    Email     string    `json:"email"`
    OtpCode   string    `json:"otp_code"`
    OtpType   string    `json:"otp_type"`
    IsUsed    bool      `json:"is_used"`
    CreatedAt time.Time `json:"created_at"`
    ExpiredAt time.Time `json:"expired_at"`
}

func NewOTPManager() OTPManager {
    return &OTP{}
}

func (otp *OTP) NewEmailVerificationOTP(email string) OTP {
    newOtp := OTP{
        Email: email,
        OtpCode: generateOTPCode(),
        OtpType: EmailVerificationOTP,
    }
    return newOtp
}

func (otp *OTP) NewPasswordChangeOTP(email string) OTP {
    newOtp := OTP{
        Email: email,
        OtpCode: generateOTPCode(),
        OtpType: EmailVerificationOTP,
    }
    return newOtp
}
