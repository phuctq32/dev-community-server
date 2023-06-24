package entity

import (
	userEntity "dev_community_server/modules/user/entity"
)

type PostUpdate struct {
	Id       *string          `map:"-"`
	Author   *userEntity.User `map:"-"`
	Title    *string          `json:"title,omitempty" validate:"omitempty,min=4" map:"title"`
	Content  *string          `json:"content,omitempty" validate:"omitempty,min=4" map:"content"`
	Images   *string          `json:"images,omitempty" map:"images"`
	TopicId  *string          `json:"topic_id,omitempty" map:"topic_id"`
	TagNames []string         `json:"tag_names,omitempty" map:"-"`
	TagIds   []string         `json:"tag_ids,omitempty" map:"tag_ids"`
}
