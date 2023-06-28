package entity

import (
	"dev_community_server/common"
	"dev_community_server/modules/comment/entity"
	entity3 "dev_community_server/modules/tag/entity"
	entity2 "dev_community_server/modules/topic/entity"
	userEntity "dev_community_server/modules/user/entity"
	"go.mongodb.org/mongo-driver/bson"
)

type PostStatus uint8

const (
	Pending PostStatus = iota
	Approved
)

type Post struct {
	common.MongoId         `bson:",inline" json:",inline"`
	common.MongoTimestamps `bson:",inline" json:",inline"`
	Title                  string     `bson:"title" json:"title"`
	Content                string     `bson:"content" json:"content"`
	Images                 []string   `bson:"images" json:"images"`
	AuthorId               string     `bson:"author_id" json:"-"`
	TopicId                string     `bson:"topic_id" json:"-"`
	TagIds                 []string   `bson:"tag_ids" json:"-"`
	Status                 PostStatus `bson:"status" json:"status"`
	UpVotes                []string   `bson:"up_votes" json:"up_votes"`
	DownVotes              []string   `bson:"down_votes" json:"down_votes"`
	ViewCount              int        `bson:"view_count" json:"view_count"`
	IsBlocked              bool       `bson:"is_blocked" json:"is_blocked"`
	// Computed fields
	Score        int              `bson:"-" json:"score"`
	Author       *userEntity.User `bson:"-" json:"author,omitempty"`
	Topic        *entity2.Topic   `bson:"-" json:"topic,omitempty"`
	Tags         []entity3.Tag    `bson:"-" json:"tags"`
	CommentCount int              `bson:"-" json:"comment_count"`
	Comments     []entity.Comment `bson:"-" json:"comments,omitempty"` // Not include replies
}

func (*Post) CollectionName() string { return "posts" }

func (post *Post) MarshalBSON() ([]byte, error) {
	dataBytes, _ := bson.Marshal(post)

	var bm common.BsonMap
	if err := bson.Unmarshal(dataBytes, &bm); err != nil {
		return nil, err
	}

	if err := bm.ToObjectId("author_id"); err != nil {
		return nil, err
	}

	if err := bm.ToObjectId("topic_id"); err != nil {
		return nil, err
	}

	if err := bm.ToListObjectId("tag_ids"); err != nil {
		return nil, err
	}

	return bson.Marshal(bm)
}
