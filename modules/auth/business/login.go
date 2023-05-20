package business

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/components/jwt"
	"dev_community_server/modules/auth/entity"
)

func (biz *authBusiness) Login(ctx context.Context, data *entity.UserLogin) (*string, error) {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, err
	}

	if ok := biz.hash.ComparePassword(user.Password, data.Password); !ok {
		return nil, entity.ErrorLoginInvalid
	}

	if !user.IsVerified {
		return nil, common.NewCustomBadRequestError("User not verified")
	}

	tokenPayload := jwt.Payload{UserId: user.Id.Hex()}
	tokenStr, err := biz.jwtProvider.GenerateAccessToken(tokenPayload, biz.jwtExpiry)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return tokenStr, nil
}
