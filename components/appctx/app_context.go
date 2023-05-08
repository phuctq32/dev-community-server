package appctx

import (
	"dev_community_server/common"
	"dev_community_server/configs"
)

type AppContext interface {
	GetValidator() common.Validator
	GetSecretKey() *string
	GetAppConfig() configs.AppConfig
}

type appContext struct {
	configs   configs.AppConfig
	validator common.Validator
}

func NewAppContext(configs configs.AppConfig, validator common.Validator) *appContext {
	return &appContext{configs: configs, validator: validator}
}

func (appCtx *appContext) GetAppConfig() configs.AppConfig {
	return appCtx.configs
}

func (appCtx *appContext) GetValidator() common.Validator {
	return appCtx.validator
}

func (appCtx *appContext) GetSecretKey() *string {
	return appCtx.configs.GetSecretKey()
}
