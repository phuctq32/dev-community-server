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
	FirstName          string     `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName           string     `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email              string     `bson:"email,omitempty" json:"email,omitempty"`
	Password           string     `bson:"password,omitempty" json:"-"`
	Birthday           *time.Time `bson:"birthday,omitempty" json:"birthday,omitempty"`
	VerifiedToken      *Token     `bson:"verified_token,omitempty" json:"-"`
	IsVerified         bool       `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
}

func NewUser(user *UserCreate) *User {
	now := time.Now()
	birthday := time.Time(user.Birthday)
	return &User{
		ModelCommon: common.ModelCommon{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		VerifiedToken: &Token{
			Token:     user.VerifiedToken,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour * 24 * 7)),
		},
		IsVerified: false,
		Birthday:   &birthday,
	}
}

func (User) CollectionName() string { return "users" }

func (u *User) GetUserId() string {
	return u.Id.Hex()
}
