package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
)

func (repo *postRepository) Find(ctx context.Context, filter entity.Filter) ([]*entity.Post, error) {
	if id, ok := filter.Other["author_id"]; ok {
		objId, err := primitive.ObjectIDFromHex(id.(string))
		if err != nil {
			return nil, err
		}

		filter.Other["author_id"] = objId
	}

	var pipeline []bson.M
	// Paginantion
	if filter.Page != nil && filter.Limit != nil {
		aggSkip := bson.M{"$skip": int64(math.Abs(float64(*filter.Page-1))) * int64(*filter.Limit)}
		aggLimit := bson.M{"$limit": int64(*filter.Limit)}
		pipeline = append(pipeline, aggSkip, aggLimit)
	}

	// Other condition
	if len(filter.Other) > 0 {
		var andStage []bson.M
		for i, condition := range filter.Other {
			andStage = append(andStage, bson.M{i: condition})
		}
		pipeline = append(pipeline, bson.M{"$match": bson.M{"$and": andStage}})
	}

	// Populate author
	userPipeline := []bson.M{
		{"$project": bson.M{
			"_id":        1,
			"first_name": 1,
			"last_name":  1,
		}},
	}
	aggPopulate := bson.M{"$lookup": bson.M{
		"from":         "users",
		"localField":   "author_id",
		"foreignField": "_id",
		"pipeline":     userPipeline,
		"as":           "author",
	}}
	aggUnwind := bson.M{"$unwind": "$author"}
	pipeline = append(pipeline, aggPopulate, aggUnwind)

	// Project
	aggProject := bson.M{"$project": bson.M{
		"author_id": 0,
	}}
	pipeline = append(pipeline, aggProject)

	cursor, err := repo.postColl.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, common.NewServerError(err)
	}

	var posts []*entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, common.NewServerError(err)
	}

	return posts, nil
}