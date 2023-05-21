package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	common.ModelCommon
	Content               string             `bson:"content" json:"content"`
	Score                 int                `bson:"score" json:"score"`
	IsAcceptedByPostOwner bool               `bson:"is_accepted_by_post_owner" json:"is_accepted_by_post_owner"`
	PostId                primitive.ObjectID `bson:"post_id" json:"post_id"`
	AuthorId              primitive.ObjectID `bson:"author_id" json:"-"`
}
