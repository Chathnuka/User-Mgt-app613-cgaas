package dao

import (
	"User-Mgt/dbConfig"
	"User-Mgt/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
    "context"
    "errors"
)

func DB_FindallUser () (*[]dto.User, error) {
	var objects []dto.User
	results, err := dbConfig.DATABASE.Collection("Users").Find(context.Background(), bson.M{ "deleted":false})
	if err != nil {
        if err == mongo.ErrNoDocuments {
        	return nil, nil
        } else {
        	return nil, err
        }
    }
	for results.Next(context.Background()) {
		var object dto.User
		if err = results.Decode(&object); err != nil {
			return nil, errors.New("Error when Decoding User")
		}
		objects = append(objects, object)
	}
	return &objects, nil
}
