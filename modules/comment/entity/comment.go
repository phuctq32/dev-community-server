package entity

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type Comment struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Content                string   `bson:"content" json:"content"`
	IsApprovedByPostAuthor bool     `bson:"is_approved_by_post_author" json:"is_accepted_by_post_owner"`
	PostId                 string   `bson:"post_id" json:"post_id"`
	ParentCommentId        *string  `bson:"parent_comment_id,omitempty" json:"parent_comment_id"`
	AuthorId               string   `bson:"author_id" json:"-"`
	UpVotes                []string `bson:"up_votes" json:"up_votes"`
	DownVotes              []string `bson:"down_votes" json:"down_votes"`
	// Computed Fields
	Author     *userEntity.User `bson:"-" json:"author"`
	Score      int              `bson:"-" json:"score"`
	Replies    []Comment        `bson:"-" json:"replies"`
	ReplyCount int              `bson:"-" json:"reply_count"`
}

func (*Comment) CollectionName() string { return "comments" }

func (cmt *Comment) MarshalBSON() ([]byte, error) {
	dataBytes, _ := bson.Marshal(cmt)

	var bm common.BsonMap
	if err := bson.Unmarshal(dataBytes, &bm); err != nil {
		return nil, err
	}

	if err := bm.ToObjectId("author_id"); err != nil {
		return nil, err
	}

	if err := bm.ToObjectId("post_id"); err != nil {
		return nil, err
	}

	if cmt.ParentCommentId != nil {
		if err := bm.ToObjectId("parent_comment_id"); err != nil {
			return nil, err
		}
	}

	return bson.Marshal(bm)
}
