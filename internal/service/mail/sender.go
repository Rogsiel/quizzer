package mail

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/rogsiel/quizzer/config"
)

type EmailSender interface {
    SendEmail(
        subject     string,
        text        string,
        to          []string,
        cc          []string,
        bcc         []string,
        attachments  []string,
    ) error
    SendWelcomeEmail(NewUserInfo) error
    SendPasswordResetEmail(ResetPasswordInfo) error
}

type Emailer struct {
    name            string
    fromAddress     string
    fromPassword    string
}

func NewEmail() EmailSender {
    return &Emailer{}
}

func (emailer *Emailer) SendEmail(
        subject     string,
        text        string,
        to          []string,
        cc          []string,
        bcc         []string,
        attachments  []string,
) error {
    config, err := config.LoadConfig("/etc/quizzer/api")
    if err != nil {
        return fmt.Errorf("Could not load config: %s", err)
    }
    
    emailer.name = config.EmailSenderName
    emailer.fromAddress = config.EmailSenderAddress
    emailer.fromPassword = config.EmailSenderPassword

    e := email.NewEmail()
    e.From = fmt.Sprintf("%s <%s>", emailer.name, emailer.fromAddress)
    e.Subject = subject
    e.HTML = []byte(text)
    e.To = to
    e.Cc = cc
    e.Bcc = bcc

    for _,attachment := range attachments {
        _, err := e.AttachFile(attachment)
        if err != nil {
            return fmt.Errorf("failed to attach file %s:%w", attachment, err)
        }
    }

    smtpAuth := smtp.PlainAuth("",
        emailer.fromAddress,
        emailer.fromPassword,
        config.SMTPAuthAddress)
    
    return e.Send(config.SMTPServerAddress, smtpAuth)
}
