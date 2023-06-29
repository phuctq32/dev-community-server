package common

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"reflect"
)

type BsonMap map[string]interface{}

func ToObjectId(val string) (primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(val)
	if err != nil {
		return primitive.NilObjectID, errors.New("invalid objectId")
	}

	return oid, nil
}

func (m BsonMap) ToMongoId(key string) error {
	if _, ok := m[key]; !ok {
		return nil
	}
	if reflect.ValueOf(m[key]).Kind() != reflect.String {
		return errors.New("type of m[key] must be string")
	}
	oid, err := ToObjectId(m[key].(string))
	if err != nil {
		return err
	}

	m[key] = oid
	return nil
}

func (m BsonMap) ToListMongoId(key string) error {
	if _, ok := m[key]; !ok {
		return nil
	}
	if reflect.ValueOf(m[key]).Kind() != reflect.Slice {
		return errors.New("type must be slice of string")
	}
	oids := make([]primitive.ObjectID, len(m[key].([]string)))
	for i, str := range m[key].([]string) {
		oid, err := ToObjectId(str)
		if err != nil {
			return err
		}
		oids[i] = oid
	}

	m[key] = oids
	return nil
}

// Define specific query for mongodb

func AppendIdQuery(filter map[string]interface{}, key string, value string) error {
	oid, err := primitive.ObjectIDFromHex(value)
	if err != nil {
		return NewServerError(errors.New("invalid objectId"))
	}
	if key == "id" {
		key = "_id"
	}
	filter[key] = oid
	return nil
}

func AppendInListQuery[T any](filter map[string]interface{}, key string, list []T) {
	if key == "id" {
		key = "_id"
	}
	filter[key] = map[string]interface{}{"$in": list}
}

func AppendInListIdQuery(filter map[string]interface{}, key string, ids []string) error {
	objectIds := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return NewServerError(errors.New("invalid objectId"))
		}
		objectIds[i] = oid
	}
	AppendInListQuery(filter, key, objectIds)
	return nil
}
