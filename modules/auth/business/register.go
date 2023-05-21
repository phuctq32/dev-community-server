package business

import (
	"context"
	"crypto/rand"
	"dev_community_server/common"
	"dev_community_server/components/mailer"
	userEntity "dev_community_server/modules/user/entity"
	"encoding/hex"
	"fmt"
)

func (biz *authBusiness) Register(ctx context.Context, data *userEntity.UserCreate) error {
	existingUser, err := biz.repo.FindOne(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return err
	}

	if existingUser != nil {
		return common.NewExistingError("user")
	}

	data.Password, err = biz.hash.HashPassword(data.Password)
	if err != nil {
		return common.NewServerError(err)
	}

	// Generate random code to verify email
	b := make([]byte, 32)
	_, err = rand.Read(b)
	if err != nil {
		return common.NewServerError(err)
	}

	verifyCode := hex.EncodeToString(b)

	mailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		*biz.appCtx.GetAppConfig().GetSendGridConfig().GetEmailFrom(),
		data.Email,
		"Verify email",
		*biz.appCtx.GetAppConfig().GetSendGridConfig().GetVerifyTemplateId(),
		map[string]interface{}{
			"username": data.FirstName + " " + data.LastName,
			"url":      fmt.Sprintf("http://localhost:3000/verification/%v", verifyCode),
		},
	)
	if err = biz.emailProvider.SendEmail(mailConfig); err != nil {
		return common.NewServerError(err)
	}
	data.VerifiedToken = verifyCode
	if err = biz.repo.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
