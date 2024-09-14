package dao

import (
	"User-Mgt/dbConfig"
	"User-Mgt/dto"
	
	"context"
    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DB_FindRolebyRoleId (roleId string) (*dto.Role, error) {
	var object dto.Role

	err := dbConfig.DATABASE.Collection("Roles").FindOne(context.Background(), bson.M{"roleid": roleId, "deleted":false}).Decode(&object)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
		    return nil, err
		}
    }
	return &object, nil
}
