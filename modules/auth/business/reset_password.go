package business

import (
	"context"
	"dev_community_server/common"
	"errors"
)

func (biz *authBusiness) ResetPassword(ctx context.Context, resetToken, newPassword string) error {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"reset_token.token": resetToken})
	if err != nil {
		return err
	}

	hashedPassword, err := biz.hash.HashPassword(newPassword)
	if err != nil {
		return common.NewServerError(errors.New("Hash not success"))
	}

	err = biz.repo.Update(ctx, user.Id.Hex(), map[string]interface{}{
		"password":    hashedPassword,
		"reset_token": nil,
	})
	return nil
}
