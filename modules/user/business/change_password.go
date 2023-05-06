package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
)

func (biz *userBusiness) ChangePassword(ctx context.Context, id string, userChangePw *entity.UserChangePassword) error {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if !biz.hash.ComparePassword(user.Password, *userChangePw.OldPassword) {
		return common.NewCustomBadRequestError("Old password incorrect")
	}

	if *userChangePw.OldPassword == *userChangePw.NewPassword {
		return common.NewCustomBadRequestError("New password must be different than old one")
	}

	hashedPassword, err := biz.hash.HashPassword(*userChangePw.NewPassword)
	if err != nil {
		return err
	}

	if err = biz.repo.Update(ctx, id, map[string]interface{}{"password": hashedPassword}); err != nil {
		return err
	}

	return nil
}
