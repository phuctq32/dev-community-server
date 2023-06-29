package entity

import (
	"dev_community_server/common"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentInsert struct {
	common.MongoId         `bson:",inline"`
	common.MongoTimestamps `bson:",inline"`
	Content                string              `bson:"content"`
	IsApprovedByPostAuthor *bool               `bson:"is_approved_by_post_author,omitempty"`
	PostId                 primitive.ObjectID  `bson:"post_id"`
	ParentCommentId        *primitive.ObjectID `bson:"parent_comment_id,omitempty"`
	AuthorId               primitive.ObjectID  `bson:"author_id"`
	UpVotes                *[]string           `bson:"up_votes,omitempty"`
	DownVotes              *[]string           `bson:"down_votes,omitempty"`
}

type Comment struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Content                string    `bson:"content" json:"content"`
	IsApprovedByPostAuthor *bool     `bson:"is_approved_by_post_author,omitempty" json:"is_accepted_by_post_owner,omitempty"`
	PostId                 string    `bson:"post_id" json:"post_id"`
	ParentCommentId        *string   `bson:"parent_comment_id,omitempty" json:"parent_comment_id"`
	AuthorId               string    `bson:"author_id" json:"-"`
	UpVotes                *[]string `bson:"up_votes,omitempty" json:"up_votes,omitempty"`
	DownVotes              *[]string `bson:"down_votes,omitempty" json:"down_votes,omitempty"`
	// Computed Fields
	Author     *userEntity.User `bson:"-" json:"author"`
	Score      *int             `bson:"-" json:"score,omitempty"`
	Replies    *[]Comment       `bson:"-" json:"replies,omitempty"`
	ReplyCount *int             `bson:"-" json:"reply_count,omitempty"`
}

func (*Comment) CollectionName() string { return "comments" }
