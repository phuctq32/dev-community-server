package configs

type SendGridConfig interface {
	GetEmailFrom() *string
	GetApiKey() *string
	GetVerifyTemplateId() *string
	GetResetPasswordTemplateId() *string
}

type sendgridConfig struct {
	SendGridApiKey          string `mapstructure:"SENDGRID_API_KEY"`
	VerifyTemplateId        string `mapstructure:"SENDGRID_VERIFY_TEMPlATE_ID"`
	ResetPasswordTemplateId string `mapstructure:"SENDGRID_RESET_PASSWORD_TEMPlATE_ID"`
	Email                   string `mapstructure:"EMAIL_FROM"`
}

func (config *sendgridConfig) GetApiKey() *string {
	return &config.SendGridApiKey
}

func (config *sendgridConfig) GetVerifyTemplateId() *string {
	return &config.VerifyTemplateId
}

func (config *sendgridConfig) GetResetPasswordTemplateId() *string {
	return &config.ResetPasswordTemplateId
}

func (config *sendgridConfig) GetEmailFrom() *string {
	return &config.Email
}
