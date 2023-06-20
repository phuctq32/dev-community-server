package entity

import (
	"dev_community_server/common"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Comment struct {
	common.ModelCommon
	Content               string              `bson:"content" json:"content"`
	Score                 int                 `bson:"score" json:"score"`
	IsAcceptedByPostOwner bool                `bson:"is_accepted_by_post_owner" json:"is_accepted_by_post_owner"`
	PostId                primitive.ObjectID  `bson:"post_id" json:"post_id"`
	ParentCommentId       *primitive.ObjectID `bson:"parent_comment_id" json:"parent_comment_id"`
	AuthorId              primitive.ObjectID  `bson:"author_id" json:"-"`
}

func NewComment(data *CommentCreate) *Comment {
	postObjId, err := primitive.ObjectIDFromHex(*data.PostId)
	authorObjId, err := primitive.ObjectIDFromHex(*data.AuthorId)
	var parentCmtObjId *primitive.ObjectID
	if data.ParentCommentId != nil {
		*parentCmtObjId, err = primitive.ObjectIDFromHex(*data.ParentCommentId)
	} else {
		parentCmtObjId = nil
	}

	if err != nil {
		panic(err)
	}

	now := time.Now()
	return &Comment{
		ModelCommon: common.ModelCommon{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Content:               *data.Content,
		Score:                 0,
		IsAcceptedByPostOwner: false,
		PostId:                postObjId,
		ParentCommentId:       parentCmtObjId,
		AuthorId:              authorObjId,
	}
}