package business

import (
	"context"
	"dev_community_server/common"
	"errors"
)

func (biz *userBusiness) ChangePassword(ctx context.Context, id string, newPassword string) error {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if biz.hash.ComparePassword(user.Password, newPassword) {
		return common.NewBadRequestError("New password must not match to old password", errors.New("must not match to old password"))
	}

	hashedPassword, err := biz.hash.HashPassword(newPassword)
	if err != nil {
		return err
	}

	if err = biz.repo.Update(ctx, id, map[string]interface{}{"password": hashedPassword}); err != nil {
		return err
	}

	return nil
}
