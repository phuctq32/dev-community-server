package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
	"time"
)

func (biz *userBusiness) UpdateUser(ctx context.Context, id string, updatingUser *entity.UserUpdate) error {
	data, err := common.StructToMap(updatingUser)
	if err != nil {
		return common.NewServerError(err)
	}

	//if val, ok := data["birthday"]; ok {
	//	data["birthday"] = time.Time(val.(common.Date))
	//}
	if updatingUser.Birthday != nil {
		data["birthday"] = time.Time(*updatingUser.Birthday)
	}

	if err = biz.userRepo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
