package dao

import (
	"User-Mgt/dbConfig"
	"User-Mgt/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindUserbyUserId (userId string) (*dto.User, error) {
	var object dto.User

	err := dbConfig.DATABASE.Collection("Users").FindOne(context.Background(), bson.M{"userid": userId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
