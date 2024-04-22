package mail

import (
	"fmt"

	"github.com/rogsiel/quizzer/config"
)

var(
    subject string = "Welcome to Quizzer"
    text    string = "Hi %s ! Please click the link to  verify : %s"
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
        "%sverify_email/%s/%s/email_verification",
        config.OriginHost,
        userInfo.Email,
        userInfo.OtpCode,
        )
    
    text = fmt.Sprintf(text, userInfo.UserName, verificationLink)

    err = emailer.SendEmail(subject, text, []string{userInfo.Email}, nil, nil, nil)
    if err != nil {
        return err
    }
    return nil
}
