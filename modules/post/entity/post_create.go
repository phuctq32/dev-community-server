package entity

import userEntity "dev_community_server/modules/user/entity"

type PostCreate struct {
	Title    string    `json:"title" validate:"required"`
	Content  string    `json:"content" validate:"required"`
	Images   []string  `json:"images,omitempty"`
	TopicId  string    `json:"topic_id" validate:"required,mongodb"`
	TagNames []*string `json:"tag_names,omitempty"`
	TagIds   []*string
	Author   *userEntity.User
}
