package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Token struct {
	Token     string    `bson:"token"`
	ExpiredAt time.Time `bson:"expired_at"`
}

type User struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	FirstName              string          `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName               string          `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email                  string          `bson:"email,omitempty" json:"email,omitempty"`
	Password               string          `bson:"password,omitempty" json:"-"`
	Birthday               *time.Time      `bson:"birthday,omitempty" json:"birthday,omitempty"`
	RoleId                 string          `bson:"role_id" json:"-"`
	Role                   string          `bson:"-" json:"role,omitempty"`
	RoleType               common.RoleType `bson:"-" json:"-"`
	Avatar                 string          `bson:"avatar" json:"avatar"`
	VerifiedToken          *Token          `bson:"verified_token,omitempty" json:"-"`
	IsVerified             bool            `bson:"is_verified,omitempty" json:"is_verified,omitempty"`
	SavedPostIds           []string        `bson:"saved_post_ids,omitempty" json:"-"`
}

func (*User) CollectionName() string { return "users" }

func (u *User) GetUserId() string {
	return *u.Id
}

func (u *User) GetRoleType() common.RoleType {
	return u.RoleType
}

func (u *User) MarshalBSON() ([]byte, error) {
	dataBytes, _ := bson.Marshal(u)
	var bm common.BsonMap
	if err := bson.Unmarshal(dataBytes, &bm); err != nil {
		return nil, err
	}
	if err := bm.ToObjectId("role_id"); err != nil {
		return nil, err
	}

	return bson.Marshal(bm)
}
