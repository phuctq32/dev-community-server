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
	user, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return err
	}

	if user == nil {
		return common.NewNotFoundError("User", common.ErrNotFound)
	}

	b := make([]byte, 32)
	if _, err = rand.Read(b); err != nil {
		return err
	}
	resetCode := hex.EncodeToString(b)

	filter := map[string]interface{}{}
	if err = common.AppendIdQuery(filter, "id", *user.Id); err != nil {
		return err
	}
	if _, err = biz.userRepo.Update(ctx, filter, map[string]interface{}{
		"reset_token": &entity.Token{
			Token:     resetCode,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour)),
		},
	}); err != nil {
		return err
	}

	mailConfig := mailer.NewEmailConfigWithDynamicTemplate(
		*biz.appCtx.GetSendGridConfig().GetEmailFrom(),
		user.Email,
		"Reset password",
		*biz.appCtx.GetSendGridConfig().GetResetPasswordTemplateId(),
		map[string]interface{}{
			"url": fmt.Sprintf("http://localhost:3000/reset-password/%v", resetCode),
		},
	)
	err = biz.emailProvider.SendEmail(mailConfig)
	if err != nil {
		return common.NewServerError(errors.New("Send mail failed"))
	}

	return nil
}
