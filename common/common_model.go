package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
	"time"
)

type MongoId struct {
	Id primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
}

type MongoTimestamps struct {
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}

// ConvertFieldToObjectId Use for map
func ConvertFieldToObjectId(m map[string]interface{}, fieldNames map[string]string) error {
	for fieldName, targetFieldName := range fieldNames {
		if value, ok := m[fieldName]; ok {
			var finalValue any
			if reflect.ValueOf(value).Kind() == reflect.Slice {
				res := make([]primitive.ObjectID, len(value.([]string)))
				for i, v := range value.([]string) {
					objectId, err := primitive.ObjectIDFromHex(v)
					res[i] = objectId
					if err != nil {
						return err
					}
				}
				finalValue = res
			} else {
				objectId, err := primitive.ObjectIDFromHex(value.(string))
				if err != nil {
					return err
				}
				finalValue = objectId
			}

			delete(m, fieldName)
			m[targetFieldName] = finalValue
		}
	}
	return nil
}
