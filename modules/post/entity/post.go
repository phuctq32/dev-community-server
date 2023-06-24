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
	PENDING PostStatus = iota
	APPROVED
)

type Post struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Title                  string               `bson:"title" json:"title"`
	Content                string               `bson:"content" json:"content"`
	Images                 []string             `bson:"images,omitempty" json:"images,omitempty"`
	AuthorId               *primitive.ObjectID  `bson:"author_id" json:"-"`
	TopicId                *primitive.ObjectID  `bson:"topic_id" json:"-"`
	TagIds                 []primitive.ObjectID `bson:"tag_ids" json:"-"`
	Status                 PostStatus           `bson:"status" json:"status"`
	UpVotes                []primitive.ObjectID `bson:"up_votes" json:"up_votes"`
	DownVotes              []primitive.ObjectID `bson:"down_votes" json:"down_votes"`
	ViewCount              int                  `bson:"view_count" json:"view_count"`
	IsBlocked              bool                 `bson:"is_blocked" json:"is_blocked"`
	// Computed fields
	Score        int               `bson:"-" json:"score"`
	Author       *userEntity.User  `bson:"-" json:"author,omitempty"`
	Topic        *entity2.Topic    `bson:"-" json:"topic,omitempty"`
	Tags         *[]entity3.Tag    `bson:"-" json:"tags,omitempty"`
	CommentCount int               `bson:"-" json:"comment_count"`
	Comments     *[]entity.Comment `bson:"-" json:"comments,omitempty"` // Not include replies
}

func NewPost(data *PostCreate) *Post {
	topicObjectId, _ := primitive.ObjectIDFromHex(data.TopicId)
	tagObjectIds := make([]primitive.ObjectID, len(data.TagIds))
	for i, tagId := range data.TagIds {
		tagObjectId, _ := primitive.ObjectIDFromHex(tagId)
		tagObjectIds[i] = tagObjectId
	}
	now := time.Now()
	//emptyArrayComments := make([]entity.Comment, 0)
	return &Post{
		MongoTimestamps: common.MongoTimestamps{
			CreatedAt: &now,
			UpdatedAt: &now,
		},
		Title:     data.Title,
		Content:   data.Content,
		AuthorId:  &data.Author.Id,
		Images:    data.Images,
		TopicId:   &topicObjectId,
		TagIds:    tagObjectIds,
		Status:    PENDING,
		UpVotes:   []primitive.ObjectID{},
		DownVotes: []primitive.ObjectID{},
		ViewCount: 0,
		IsBlocked: false,
		// Init computed fields
		Score:        0,
		CommentCount: 0,
		Comments:     &[]entity.Comment{},
		Author:       data.Author,
	}
}
