package sgmail

import (
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendSimpleMail(params *Params) error {
	m := mail.NewV3Mail()
	m.SetFrom(params.GetFrom())
	m.Subject = params.Subject
	m.AddContent(params.GetContent())
	if params.HasCategory() {
		m.AddCategories(params.Category)
	}
	return nil
}
