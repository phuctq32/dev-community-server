package configs

type SendGridConfigs interface {
	GetEmailFrom() *string
	GetApiKey() *string
	GetVerifyTemplateId() *string
	GetResetPasswordTemplateId() *string
}

type sgConfigs struct {
	SendGridApiKey          string `mapstructure:"SENDGRID_API_KEY"`
	VerifyTemplateId        string `mapstructure:"SENDGRID_VERIFY_TEMPlATE_ID"`
	ResetPasswordTemplateId string `mapstructure:"SENDGRID_RESET_PASSWORD_TEMPlATE_ID"`
}

func (configs *appConfigs) GetApiKey() *string {
	return &configs.SendGridApiKey
}

func (configs *appConfigs) GetVerifyTemplateId() *string {
	return &configs.VerifyTemplateId
}

func (configs *appConfigs) GetResetPasswordTemplateId() *string {
	return &configs.ResetPasswordTemplateId
}

func (configs *appConfigs) GetEmailFrom() *string {
	return &configs.Email
}
