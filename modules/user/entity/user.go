package entity

import (
	"dev_community_server/common"
	"time"
)

type User struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Name               string    `bson:"name" json:"name"`
	Email              string    `bson:"email" json:"email"`
	Password           string    `bson:"password" json:"-"`
	Birthday           time.Time `bson:"birthday" json:"birthday"`
	IsVerified         bool      `bson:"is_verified" json:"is_verified"`
}

func NewUser(user *UserCreate) *User {
	return &User{
		ModelCommon: common.ModelCommon{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Email:      user.Email,
		Name:       user.FirstName + " " + user.LastName,
		Password:   user.Password,
		IsVerified: false,
		Birthday:   time.Time(user.Birthday),
	}
}

func (User) CollectionName() string { return "users" }
