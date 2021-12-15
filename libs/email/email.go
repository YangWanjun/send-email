package email

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"send-email/model/entity"
)

func send(emailConfig *entity.EmailConfigEntity) {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetAddressHeader("From", emailConfig.GetSender(), "NoRelay")
	m.SetAddressHeader("Sender", emailConfig.GetSender(), emailConfig.GetDisplayName())
	m.SetHeader("Reply-To", emailConfig.GetSender())

	// Set E-Mail receivers
	m.SetAddressHeader("To", "", "")

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", "This is Gomail test body")

	// Settings for SMTP server
	d := gomail.NewDialer(emailConfig.SmtpServer, emailConfig.SmtpPort, emailConfig.Username, emailConfig.Password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: "e-business.co.jp"}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return
}
