package configs

type SendGridConfigs interface {
	GetApiKey() *string
	GetVerifyTemplateId() *string
}

type sgConfigs struct {
	SendGridApiKey   string `mapstructure:"SENDGRID_API_KEY"`
	VerifyTemplateId string `mapstructure:"SENDGRID_VERIFY_TEMPlATE_ID"`
}

func (configs *appConfigs) GetApiKey() *string {
	return &configs.SendGridApiKey
}

func (configs *appConfigs) GetVerifyTemplateId() *string {
	return &configs.VerifyTemplateId
}
