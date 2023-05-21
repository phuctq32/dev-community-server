package entity

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Title              string              `bson:"title" json:"title"`
	Content            string              `bson:"content" json:"content"`
	Images             []string            `bson:"images,omitempty" json:"images,omitempty"`
	AuthorId           *primitive.ObjectID `bson:"author_id" json:"author_id,omitempty"`
	Author             *userEntity.User    `bson:"author,omitempty" json:"author"`
}

func NewPost(data *PostCreate) *Post {
	authorObjId, _ := primitive.ObjectIDFromHex(data.AuthorId)
	now := time.Now()
	return &Post{
		ModelCommon: common.ModelCommon{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Title:    data.Title,
		Content:  data.Content,
		AuthorId: &authorObjId,
		Images:   data.Images,
	}
}
