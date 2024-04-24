package mail

import (
	"fmt"

	"github.com/rogsiel/quizzer/config"
)

var(
    welcomSubject string = "Welcome to Quizzer"
    welcomText    string = "Hi %s ! Please click the link to  verify : %s"
)
type NewUserInfo struct{
    UserName    string
    Email       string
    OtpCode     string
}
func (emailer *Emailer) SendWelcomeEmail(userInfo NewUserInfo) error {
    config, err := config.LoadConfig("/etc/quizzer/api")
    if err != nil {
        return fmt.Errorf("Could not load config: %s", err)
    }

    verificationLink := fmt.Sprintf(
        "%sverify_email?otp_type=email_verification&otp_code=%s",
        config.OriginHost,
        userInfo.OtpCode,
        )
    
    welcomText = fmt.Sprintf(welcomText, userInfo.UserName, verificationLink)

    err = emailer.SendEmail(
        welcomSubject,
        welcomText,
        []string{userInfo.Email},
        nil, nil, nil,
        )
    if err != nil {
        return err
    }
    return nil
}
