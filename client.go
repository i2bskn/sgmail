package sgmail

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var MailClient *Client

type Client struct {
	SendGridClient *sendgrid.Client
}

func NewClient(apiKey string) *Client {
	client := &Client{
		SendGridClient: sendgrid.NewSendClient(apiKey),
	}
	return client
}

func (client *Client) Send(params *Params) error {
	m := mail.NewV3Mail()
	m.SetFrom(params.GetFrom())
	m.Subject = params.Subject
	m.AddContent(params.GetContent())
	if params.HasCategory() {
		m.AddCategories(params.Category)
	}
	client.SendGridClient.Send(m)
	return nil
}
