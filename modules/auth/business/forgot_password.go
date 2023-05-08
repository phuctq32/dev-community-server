package business

import (
	"context"
	"crypto/rand"
	"dev_community_server/common"
	"dev_community_server/components/mailer"
	"dev_community_server/modules/user/entity"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

func (biz *authBusiness) ForgotPassword(ctx context.Context, email string) error {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}

	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return err
	}
	resetCode := hex.EncodeToString(b)

	err = biz.repo.Update(ctx, user.Id.Hex(), map[string]interface{}{
		"reset_token": &entity.Token{
			Token:     resetCode,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour)),
		},
	})

	mailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		*biz.appCtx.GetAppConfig().GetSendGridConfig().GetEmailFrom(),
		user.Email,
		"Reset password",
		*biz.appCtx.GetAppConfig().GetSendGridConfig().GetResetPasswordTemplateId(),
		map[string]interface{}{
			"url": fmt.Sprintf("http://localhost:8080/api/v1/auth/reset_password/%v", resetCode),
		},
	)
	err = biz.emailProvider.SendEmail(mailConfig)
	if err != nil {
		return common.NewServerError(errors.New("Send mail failed"))
	}

	return nil
}
