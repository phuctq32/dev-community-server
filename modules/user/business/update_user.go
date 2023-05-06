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

	if val, ok := data["birthday"]; ok {
		data["birthday"] = time.Time(val.(common.Date))
	}

	if err = biz.repo.Update(ctx, id, data); err != nil {
		return err
	}

	return nil
}
