package interfaces

import (
	gomail "gopkg.in/gomail.v2"
)

// IMail interface
type IMail interface {
	GetInstance() *gomail.Dialer
	SendMail(body, subject, to string) error
}
