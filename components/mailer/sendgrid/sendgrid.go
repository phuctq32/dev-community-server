package sendgrid

import (
	"dev_community_server/components/mailer"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
)

type sendGridService struct {
	ApiKey string
}

func NewSendGridService(key string) *sendGridService {
	return &sendGridService{ApiKey: key}
}

func (sg *sendGridService) SendEmail(config *mailer.EmailConfig) error {
	client := sendgrid.NewSendClient(sg.ApiKey)

	fromEmail := mail.NewEmail("Dev Community", config.From)
	toEmail := mail.NewEmail("", config.To)
	mailData := mail.NewV3Mail()
	mailData.SetFrom(fromEmail)
	mailData.SetTemplateID(config.TemplateId)

	personalization := mail.NewPersonalization()
	personalization.AddTos(toEmail)
	mailData.AddPersonalizations(personalization)
	config.Data["subject"] = config.Subject
	personalization.DynamicTemplateData = config.Data

	res, err := client.Send(mailData)
	log.Printf("sent mail%+v\n", res)
	if err != nil {
		return err
	}

	return nil
}
