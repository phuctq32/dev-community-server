package entity

import (
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	entity3 "dev_community_server/modules/tag/entity"
	entity2 "dev_community_server/modules/topic/entity"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PostStatus uint8

const (
	Pending PostStatus = iota
	Approved
)

type PostInsert struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Title                  string               `bson:"title"`
	Content                string               `bson:"content"`
	Images                 []string             `bson:"images"`
	AuthorId               primitive.ObjectID   `bson:"author_id"`
	TopicId                primitive.ObjectID   `bson:"topic_id"`
	TagIds                 []primitive.ObjectID `bson:"tag_ids"`
	Status                 PostStatus           `bson:"status"`
	UpVotes                []primitive.ObjectID `bson:"up_votes"`
	DownVotes              []primitive.ObjectID `bson:"down_votes"`
	ViewCount              int                  `bson:"view_count"`
	IsBlocked              bool                 `bson:"is_blocked"`
}

type Post struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Title                  string      `bson:"title" json:"title"`
	Content                string      `bson:"content" json:"content"`
	Images                 []string    `bson:"images" json:"images"`
	AuthorId               string      `bson:"author_id" json:"-"`
	TopicId                string      `bson:"topic_id" json:"-"`
	TagIds                 []string    `bson:"tag_ids" json:"-"`
	Status                 PostStatus  `bson:"status" json:"status"`
	UpVotes                []string    `bson:"up_votes" json:"up_votes"`
	DownVotes              []string    `bson:"down_votes" json:"down_votes"`
	TimestampViews         []time.Time `bson:"timestamp_views" json:"-"`
	IsBlocked              bool        `bson:"is_blocked" json:"is_blocked"`
	// Computed fields
	ViewsPerThreeDays int              `bson:"-" json:"-"`
	ViewCount         int              `bson:"-" json:"view_count"`
	Score             int              `bson:"-" json:"score"`
	Author            *userEntity.User `bson:"-" json:"author,omitempty"`
	Topic             *entity2.Topic   `bson:"-" json:"topic,omitempty"`
	Tags              []entity3.Tag    `bson:"-" json:"tags"`
	CommentCount      int              `bson:"-" json:"comment_count"`
	Comments          []entity.Comment `bson:"-" json:"comments,omitempty"` // Not include replies
}

func (*Post) CollectionName() string { return "posts" }
