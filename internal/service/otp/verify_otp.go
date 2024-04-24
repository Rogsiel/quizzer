package otp

import (
	"errors"
	"time"
)

func (otp *OTP) VerifyOTP(otpType string) error {
    if otp.OtpType != otpType {
        return errors.New("Invalid security code")
    }
    if !otp.ExpiredAt.After(time.Now()) {
        return errors.New("Security Code has expired")
    }
    if otp.IsUsed {
        return errors.New("Invalid Security Code")
    }
    return nil
}
