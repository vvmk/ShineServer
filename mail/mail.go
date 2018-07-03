// Package Mail is an abstraction around the gomail package that handles
// sending user email confirmation messages.
package mail

import (
	"os"

	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendConfirmation builds a message and dispatches a confirmation email
// to the supplied address.
func SendConfirmation(tag string, address string) (*rest.Response, error) {
	from := mail.NewEmail("vv", "ssrbot@ssr.complexaesthetic.com")
	subject := "Welcome to the party!"
	to := mail.NewEmail(tag, address)

	plainTextContent := "Thanks for signing up and jumping through this extra hoop. You can confirm your email using the link below:\n"
	htmlContent := "Thanks for signing up and jumping through this extra hoop. You can confirm your email using the link below:\n"

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
