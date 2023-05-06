package business

import (
	"context"
	"dev_community_server/common"
	"time"
)

func (biz *authBusiness) VerifyEmail(ctx context.Context, verifyToken string) error {
	user, err := biz.repo.FindOne(ctx, map[string]interface{}{"verified_token.token": verifyToken})
	if err != nil {
		return err
	}

	if user.VerifiedToken.ExpiredAt.Before(time.Now()) {
		return common.NewCustomBadRequestError("Verification code expired")
	}

	if err = biz.repo.Update(ctx, user.Id.Hex(), map[string]interface{}{
		"is_verified":    true,
		"verified_token": nil,
	}); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
