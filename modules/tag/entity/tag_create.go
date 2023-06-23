package entity

type TagCreate struct {
	Name    string `json:"name" validate:"required" map:"name"`
	TopicId string `json:"topic_id" validate:"required,mongodb" map:"topic_id" toObjectId:"true"`
}
