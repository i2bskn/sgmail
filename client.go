package sgmail

import (
	"github.com/sendgrid/sendgrid-go"
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

func (client *Client) Send() error {
	return nil
}
