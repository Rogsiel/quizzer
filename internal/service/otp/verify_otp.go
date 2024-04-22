package otp

import (
	"errors"
	"time"
)

func (otp *OTP) VerifyOTP() error {
    if otp.OtpType != EmailVerificationOTP {
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
