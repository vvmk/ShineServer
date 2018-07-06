// Package Mail is an abstraction around the gomail package that handles
// sending user email confirmation messages.
package mail

import (
	"fmt"
	"os"

	"github.com/sendgrid/rest"
	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type MessageData struct {
	Address string
	Tag     string
	UserId  int
	Token   string
}

// SendConfirmation builds and dispatches an email confirmation to new users.
func SendConfirmation(data *MessageData) (*rest.Response, error) {
	from := mail.NewEmail("vvmk", "vv@shfflshinerepeat.com")
	subject := "Welcome to the party!"
	to := mail.NewEmail(data.Tag, data.Address)
	confirm := fmt.Sprintf("https://shfflshinerepeat.com/confirm?uid=%d&token=%s", data.UserId, data.Token)

	message := new(mail.SGMailV3)
	message.SetTemplateID(os.Getenv("SENDGRID_TEMPLATE_CONFIRMATION"))
	message.SetFrom(from)
	message.Subject = subject

	p := mail.NewPersonalization()
	p.AddTos(to)
	p.SetSubstitution("-confirmlink-", confirm)
	p.SetSubstitution("-ssrtag-", data.Tag)
	message.AddPersonalizations(p)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}
