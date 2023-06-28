package business

import (
	"context"
	"dev_community_server/common"
	"errors"
)

func (biz *authBusiness) ResetPassword(ctx context.Context, resetToken, newPassword string) error {
	user, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"reset_token.token": resetToken})
	if err != nil {
		return err
	}

	if user == nil {
		return common.NewNotFoundError("User", common.ErrNotFound)
	}

	hashedPassword, err := biz.hash.HashPassword(newPassword)
	if err != nil {
		return common.NewServerError(errors.New("Hash not success"))
	}

	filter := map[string]interface{}{}
	if err = common.AppendIdQuery(filter, "id", *user.Id); err != nil {
		return err
	}
	if _, err = biz.userRepo.Update(ctx, filter, map[string]interface{}{
		"password":    hashedPassword,
		"reset_token": nil,
	}); err != nil {
		return err
	}
	return nil
}
