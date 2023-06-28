package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
)

func (biz *userBusiness) ChangePassword(ctx context.Context, id string, userChangePw *entity.UserChangePassword) error {
	filter := map[string]interface{}{}
	if err := common.AppendIdQuery(filter, "id", id); err != nil {
		return err
	}
	user, err := biz.userRepo.FindOne(ctx, filter)
	if err != nil {
		return err
	}

	if user == nil {
		return common.NewNotFoundError("User", common.ErrNotFound)
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

	if _, err = biz.userRepo.Update(ctx, filter, map[string]interface{}{"password": hashedPassword}); err != nil {
		return err
	}

	return nil
}
