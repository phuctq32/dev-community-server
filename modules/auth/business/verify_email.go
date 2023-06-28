package business

import (
	"context"
	"dev_community_server/common"
	"time"
)

func (biz *authBusiness) VerifyEmail(ctx context.Context, verifyToken string) error {
	user, err := biz.userRepo.FindOne(ctx, map[string]interface{}{"verified_token.token": verifyToken})
	if err != nil {
		return err
	}

	if user == nil {
		return common.NewCustomBadRequestError("Invalid token")
	}

	if user.VerifiedToken.ExpiredAt.Before(time.Now()) {
		return common.NewCustomBadRequestError("Verification token expired")
	}

	filter := map[string]interface{}{}
	if err = common.AppendIdQuery(filter, "id", *user.Id); err != nil {
		return err
	}
	if _, err = biz.userRepo.Update(ctx, filter, map[string]interface{}{
		"is_verified":    true,
		"verified_token": nil,
	}); err != nil {
		return common.NewServerError(err)
	}

	return nil
}
