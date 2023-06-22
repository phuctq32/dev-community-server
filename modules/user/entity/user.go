package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Token struct {
	Token     string    `bson:"token"`
	ExpiredAt time.Time `bson:"expired_at"`
}

type User struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	FirstName              string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName               string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email                  string             `bson:"email,omitempty" json:"email,omitempty"`
	Password               string             `bson:"password,omitempty" json:"-"`
	Birthday               *time.Time         `bson:"birthday,omitempty" json:"birthday,omitempty"`
	RoleId                 primitive.ObjectID `bson:"role_id" json:"-"`
	Role                   string             `json:"role,omitempty"`
	RoleType               common.RoleType    `json:"-"`
	Avatar                 string             `bson:"avatar" json:"avatar"`
	VerifiedToken          *Token             `bson:"verified_token,omitempty" json:"-"`
	IsVerified             bool               `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
}

func NewUser(user *UserCreate) *User {
	now := time.Now()
	birthday := time.Time(user.Birthday)
	return &User{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Birthday:  &birthday,
		Avatar:    common.DEFAULT_AVATAR_URL,
		RoleId:    user.RoleId,
		VerifiedToken: &Token{
			Token:     user.VerifiedToken,
			ExpiredAt: time.Now().Add(time.Duration(time.Hour * 24 * 7)),
		},
		IsVerified: false,
	}
}

func (User) CollectionName() string { return "users" }

func (u *User) GetUserId() string {
	return u.Id.Hex()
}

func (u *User) GetRoleType() common.RoleType {
	return u.RoleType
}
