package configs

type CloudinaryConfig interface {
	GetApiKey() *string
	GetApiSecret() *string
	GetCloudName() *string
}

type cloudinaryConfig struct {
	ApiKey    string `mapstructure:"222345215317385"`
	ApiSecret string `mapstructure:"CLOUDINARY_API_SECRET"`
	CloudName string `mapstructure:"CLOUDINARY_CLOUD_NAME"`
}

func (config *cloudinaryConfig) GetApiKey() *string {
	return &config.ApiKey
}

func (config *cloudinaryConfig) GetApiSecret() *string {
	return &config.ApiSecret
}

func (config *cloudinaryConfig) GetCloudName() *string {
	return &config.CloudName
}
