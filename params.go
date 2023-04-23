package sgmail

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

const (
	DefaultContentType = "text/plain"
)

type Params struct {
	Subject     string
	ContentType string
	Content     string
	Category    string
	FromName    string
	FromEmail   string
	ToName      string
	ToEmail     string
}

func NewParams() *Params {
	params := &Params{
		ContentType: DefaultContentType,
	}
	return params
}

func (params *Params) GetContentType() string {
	if len(params.ContentType) == 0 {
		return DefaultContentType
	}

	return params.ContentType
}

func (params *Params) GetFrom() *mail.Email {
	return mail.NewEmail(params.FromName, params.FromEmail)
}

func (params *Params) GetContent() *mail.Content {
	return mail.NewContent(params.GetContentType(), params.Content)
}

func (params *Params) HasCategory() bool {
	return len(params.Category) > 0
}
