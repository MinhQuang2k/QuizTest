package mail

import (
	"quiztest/app/interfaces"
	"quiztest/config"
	"quiztest/pkg/logger"

	gomail "gopkg.in/gomail.v2"
)

type emailMgr struct {
	mail *gomail.Dialer
}

func NewMail() interfaces.IMail {
	cfg := config.GetConfig()

	// Sender data.
	smtpPass := cfg.SmtpPass
	smtpUser := cfg.SmtpUser
	smtpHost := cfg.SmtpHost
	smtpPort := cfg.SmtpPort

	mail := gomail.NewPlainDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	logger.Info("ಠ‿ಠ Gomail ad running ಠ‿ಠ")
	return &emailMgr{
		mail: mail,
	}
}

// GetInstance get mail instance
func (ml *emailMgr) GetInstance() *gomail.Dialer {
	return ml.mail
}

func (ml *emailMgr) SendMail(body, subject, to string) error {
	cfg := config.GetConfig()
	from := cfg.EmailFrom
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if err := ml.mail.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
