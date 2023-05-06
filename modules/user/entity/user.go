package entity

import (
	"dev_community_server/common"
	"time"
)

type Token struct {
	Token     string    `bson:"token"`
	ExpiredAt time.Time `bson:"expired_at"`
}

type User struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Name               string    `bson:"name" json:"name"`
	Email              string    `bson:"email" json:"email"`
	Password           string    `bson:"password" json:"-"`
	Birthday           time.Time `bson:"birthday" json:"birthday"`
	VerifiedToken      *Token    `bson:"verified_token" json:"-"`
	IsVerified         bool      `bson:"is_verified" json:"is_verified"`
}

func NewUser(user *UserCreate) *User {
	return &User{
		ModelCommon: common.ModelCommon{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:    user.Email,
		Name:     user.FirstName + " " + user.LastName,
		Password: user.Password,
		VerifiedToken: &Token{
			Token:     user.VerifiedToken,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour * 24 * 7)),
		},
		IsVerified: false,
		Birthday:   time.Time(user.Birthday),
	}
}

func (User) CollectionName() string { return "users" }
