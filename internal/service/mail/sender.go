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
}

type Emailer struct {
    name            string
    fromAddress     string
    fromPassword    string
}

func NewEmail(name string, fromAddress string, fromPassword string) EmailSender {
    return &Emailer{
        name: name,
        fromAddress: fromAddress,
        fromPassword: fromPassword,
    }
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
