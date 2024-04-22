package model

import "time"

type OTP struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	OtpCode   string    `json:"otp_code"`
	OtpType   string    `json:"otp_type"`
	IsUsed    bool      `json:"is_used"`
	CreatedAt time.Time `json:"created_at"`
	ExpiredAt time.Time `json:"expired_at"`
}
