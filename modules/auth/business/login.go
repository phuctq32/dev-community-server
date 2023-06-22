package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/jwt"
	"dev_community_server/modules/auth/entity"
	userEntity "dev_community_server/modules/user/entity"
)

func (biz *authBusiness) Login(ctx context.Context, data *entity.UserLogin) (*string, *userEntity.User, error) {
	user, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, nil, err
	}

	if user == nil {
		return nil, nil, common.NewNotFoundError("User", common.ErrNotFound)
	}

	if ok := biz.hash.ComparePassword(user.Password, data.Password); !ok {
		return nil, nil, entity.ErrorLoginInvalid
	}

	if !user.IsVerified {
		return nil, nil, common.NewCustomBadRequestError("User not verified")
	}

	tokenPayload := jwt.Payload{UserId: user.Id.Hex()}
	tokenStr, err := biz.jwtProvider.GenerateAccessToken(tokenPayload, biz.jwtExpiry)
	if err != nil {
		return nil, nil, err
	}

	if err != nil {
		return nil, nil, err
	}

	return tokenStr, user, nil
}
