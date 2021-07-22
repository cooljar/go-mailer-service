package usecase

import (
	"bytes"
	"fmt"
	"github.com/cooljar/go-mailer-service/domain"
	"gopkg.in/gomail.v2"
	"html/template"
	"net"
	"net/smtp"
	"os"
	"path"
	"strings"
	"time"
)

const forceDisconnectAfter = time.Second * 5

type mailerUsecase struct {
	dialer *gomail.Dialer
}

func NewMailerUsecase(dialer *gomail.Dialer) domain.MailerUseCase {
	return &mailerUsecase{dialer: dialer}
}

func (m mailerUsecase) SendMail(mail domain.Mail) error {
	var pathLayout = path.Join("assets", "html", "layout", "mail.html")
	var pathContent = path.Join("assets", "html", "email-content.html")
	var dataEmail = map[string]interface{}{
		"subject":mail.Subject,
		"content":  mail.Body,
	}

	var tpl bytes.Buffer

	templateEmail, err := template.ParseFiles(pathLayout, pathContent)
	if err != nil {
		return err
	}

	err = templateEmail.Execute(&tpl, dataEmail)
	if err != nil {
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", mail.SenderName + " <" + os.Getenv("SMTP_EMAIL_ADDRESS") + ">")
	mailer.SetHeader("To", mail.To)
	//mailer.SetAddressHeader("Cc", "tralalala@gmail.com", "Tra Lala La")
	mailer.SetHeader("Subject", mail.Subject)
	mailer.SetBody("text/html", tpl.String())
	//mailer.Attach("./sample.png")

	err = m.dialer.DialAndSend(mailer)
	if err != nil {
		return err
	}

	return nil
}

// IsEmailAddressValid validate mail host.
func isEmailAddressValid(email string) (valid bool, err error) {
	_, host := split(email)
	mx, err := net.LookupMX(host)
	if err != nil {
		return false, err
	}
	client, err := dialTimeout(fmt.Sprintf("%s:%d", mx[0].Host, 25), forceDisconnectAfter)
	if err != nil {
		return false, err
	}
	client.Close()

	return true, nil
}

// DialTimeout returns a new Client connected to an SMTP server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func dialTimeout(addr string, timeout time.Duration) (*smtp.Client, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}

	t := time.AfterFunc(timeout, func() { conn.Close() })
	defer t.Stop()

	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}


func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	// If no @ present, not a valid email.
	if i < 0 {
		return
	}
	account = email[:i]
	host = email[i+1:]
	return
}
