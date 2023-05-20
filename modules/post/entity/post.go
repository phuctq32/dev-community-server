package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Title              string             `bson:"title" json:"title"`
	Content            string             `bson:"content" json:"content"`
	Images             []string           `bson:"images" json:"images"`
	AuthorId           primitive.ObjectID `bson:"author_id" json:"author_id"`
}
