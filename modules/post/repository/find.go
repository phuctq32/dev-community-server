package repository

import (
	"context"
	"dev_community_server/common"
	"dev_community_server/modules/post/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//func (repo *postRepository) Find(ctx context.Context, filter common.Filter) ([]entity.Post, error) {
//	if id, ok := filter.Other["author_id"]; ok {
//		objId, err := primitive.ObjectIDFromHex(id.(string))
//		if err != nil {
//			return nil, err
//		}
//
//		filter.Other["author_id"] = objId
//	}
//
//	var pipeline []bson.M
//
//	// Sort
//	sortStage := bson.M{"$sort": bson.M{"created_at": -1}}
//
//	if filter.Search != nil {
//		//searchTerm := common.RemoveVietnameseAccent(*filter.Search)
//		searchTerm := common.RemoveVietnameseAccent(*filter.Search)
//		log.Println(searchTerm)
//		matchTextStage := bson.M{"$match": bson.M{
//			"$text": bson.M{
//				"$search": common.RemoveVietnameseAccent(*filter.Search),
//			},
//			//"$or": []bson.M{
//			//	{"title": bson.M{"$regex": searchTerm, "$options": "i"}},
//			//	{"content": bson.M{"$regex": searchTerm, "$options": "i"}},
//			//},
//		}}
//		sortStage = bson.M{"$sort": bson.M{"$score": bson.M{"$meta": "textScore"}}}
//		pipeline = append(pipeline, matchTextStage)
//	}
//
//	// Other condition
//	//if len(filter.Other) > 0 {
//	//	var andStage []bson.M
//	//	for i, condition := range filter.Other {
//	//		andStage = append(andStage, bson.M{i: condition})
//	//	}
//	//	matchStage := bson.M{"$match": bson.M{"$and": andStage}}
//	//	pipeline = append(pipeline, matchStage)
//	//}
//
//	// Paginantion
//	if filter.Page != nil && filter.Limit != nil {
//		skipStage := bson.M{"$skip": int64(math.Abs(float64(*filter.Page-1))) * int64(*filter.Limit)}
//		limitStage := bson.M{"$limit": int64(*filter.Limit)}
//		pipeline = append(pipeline, skipStage, limitStage)
//	}
//
//	// Add sort stage
//
//	pipeline = append(pipeline, sortStage)
//
//	// Populate author
//	userPipeline := []bson.M{
//		{"$project": bson.M{
//			"_id":        1,
//			"first_name": 1,
//			"last_name":  1,
//			"avatar":     1,
//		}},
//	}
//	populateStage := bson.M{"$lookup": bson.M{
//		"from":         "users",
//		"localField":   "author_id",
//		"foreignField": "_id",
//		"pipeline":     userPipeline,
//		"as":           "author",
//	}}
//	unwindStage := bson.M{"$unwind": "$author"}
//	pipeline = append(pipeline, populateStage, unwindStage)
//
//	// Project
//	projectStage := bson.M{"$project": bson.M{
//		"author_id": 0,
//	}}
//	pipeline = append(pipeline, projectStage)
//
//	cursor, err := repo.postColl.Aggregate(ctx, pipeline)
//	if err != nil {
//		return nil, common.NewServerError(err)
//	}
//
//	var posts []entity.Post
//	if err = cursor.All(ctx, &posts); err != nil {
//		return nil, common.NewServerError(err)
//	}
//
//	return posts, nil
//}

func (repo *postRepository) Find(
	ctx context.Context,
	filter map[string]interface{},
	pagination *common.Pagination,
) ([]entity.Post, error) {
	// Sort by newest
	opts := options.Find().SetSort(bson.M{"created_at": -1})

	// Pagination
	if pagination != nil {
		opts.SetLimit(int64(*pagination.Limit)).SetSkip(int64((*pagination.Page - 1) * (*pagination.Limit)))
	}

	cursor, err := repo.postColl.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var posts []entity.Post
	if err = cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
