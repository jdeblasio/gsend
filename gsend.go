package gsend

import (
	"bytes"
	"fmt"
	"net/smtp"
)

type Gmail struct {
	From    string
	To      string
	Subject string
	Message string

	Password string
}

func (self *Gmail) Send() error {
	body := new(bytes.Buffer)
	fmt.Fprintf(body, "To: %s\r\n", self.To)
	fmt.Fprintf(body, "Subject: %s\r\n", self.Subject)
	fmt.Fprintf(body, "\r\n")
	fmt.Fprintf(body, "%s", self.Message)

	const server = "smtp.gmail.com"
	auth := smtp.PlainAuth("", self.From, self.Password, server)
	return smtp.SendMail(
		fmt.Sprintf("%s:587", server),
		auth,
		self.From,
		[]string{self.To},
		body.Bytes(),
	)
}
