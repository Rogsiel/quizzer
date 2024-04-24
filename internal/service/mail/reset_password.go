package mail

import (
	"fmt"

	"github.com/rogsiel/quizzer/config"
)

var(
    passwordSubject string = "Quizzer Reset Pssword"
    passwordText    string = "Hi %s ! Please click the link to reset password : %s"
)
type ResetPasswordInfo struct{
    UserName    string
    Email       string
    OtpCode     string
}
func (emailer *Emailer) SendPasswordResetEmail(resetInfo ResetPasswordInfo) error {
    config, err := config.LoadConfig("/etc/quizzer/api")
    if err != nil {
        return fmt.Errorf("Could not load config: %s", err)
    }

    resetPasswordLink := fmt.Sprintf(
        "%slogin/reset-password?otp_type=password_change&otp_code=%s",
        config.OriginHost,
        resetInfo.OtpCode,
        )
    
    passwordText = fmt.Sprintf(passwordText, resetInfo.UserName, resetPasswordLink)

    err = emailer.SendEmail(
        passwordSubject,
        passwordText,
        []string{resetInfo.Email},
        nil, nil, nil)
    
    if err != nil {
        return err
    }
    
    return nil
}
