package email

import (
	"crypto/tls"
	"errors"
	"regexp"

	"gopkg.in/gomail.v2"
)

func ValidateFormat(email string) error {
	emailRegexp := regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegexp.MatchString(email) {
		return errors.New("invalid format")
	}
	return nil
}

type Config struct {
	SmtpServer  string
	SmtpPort    int
	DefaultMail string
	DefaultPwd  string
}

func SendEmail(mailTo string, subject string, body string, cfg *Config, images ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(cfg.DefaultMail, cfg.DefaultMail))
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	for _, v := range images {
		m.Embed(v)
	}
	d := gomail.NewDialer(cfg.SmtpServer, cfg.SmtpPort, cfg.DefaultMail, cfg.DefaultPwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
