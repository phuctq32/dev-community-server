package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Post struct {
	common.ModelCommon `bson:",inline" json:",inline"`
	Title              string             `bson:"title" json:"title"`
	Content            string             `bson:"content" json:"content"`
	Images             []string           `bson:"images,omitempty" json:"images,omitempty"`
	AuthorId           primitive.ObjectID `bson:"author_id" json:"author_id"`
}

func NewPost(data *PostCreate) *Post {
	authorObjId, _ := primitive.ObjectIDFromHex(data.AuthorId)

	return &Post{
		ModelCommon: common.ModelCommon{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:    data.Title,
		Content:  data.Content,
		AuthorId: authorObjId,
		Images:   data.Images,
	}
}
