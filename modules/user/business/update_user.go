package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/user/entity"
	"time"
)

func (biz *userBusiness) UpdateUser(ctx context.Context, id string, data *entity.UserUpdate) (*entity.User, error) {
	updatedData, err := common.StructToMap(data)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	if data.Birthday != nil {
		updatedData["birthday"] = time.Time(*data.Birthday)
	}

	filter := map[string]interface{}{}
	if err = common.AppendIdQuery(filter, "id", id); err != nil {
		return nil, common.NewServerError(err)
	}
	user, err := biz.userRepo.Update(ctx, filter, updatedData)
	if err != nil {
		return nil, err
	}

	return user, nil
}
