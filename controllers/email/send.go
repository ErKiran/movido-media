package email

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"movido-media/controllers/billing"

	"gopkg.in/gomail.v2"
)

//go:embed email.html
var emailTemplateHTML string

func (ec emailController) Sender(ctx context.Context, data billing.ContractDetail, path string) error {
	var (
		host     = os.Getenv("SMTP_HOST")
		port     = os.Getenv("SMTP_PORT")
		username = os.Getenv("EMAIL_USERNAME")
		password = os.Getenv("EMAIL_PASSWORD")
		from     = os.Getenv("EMAIL_FROM")
	)
	p, _ := strconv.Atoi(port)

	dialer := gomail.NewDialer(host, p, username, password)

	ec.mailer.SetHeader("From", from)
	ec.mailer.SetHeader("To", data.Email)

	// Parse the email template
	templateHTML := template.Must(template.New("email").Parse(emailTemplateHTML))

	// Create a buffer to store the rendered template
	var renderedTemplate bytes.Buffer

	// Execute the template with the data and write the result to the buffer
	err := templateHTML.Execute(&renderedTemplate, data)
	if err != nil {
		return err
	}

	// Set the HTML body
	ec.mailer.SetBody("text/html", renderedTemplate.String())

	ec.mailer.SetHeader("Subject", fmt.Sprintf("Invoice for %s", data.ProductName))
	ec.mailer.Attach(path)

	if err := dialer.DialAndSend(ec.mailer); err != nil {
		return err
	}

	return nil
}
