package dao

import (
	"User-Mgt/dbConfig"
    "User-Mgt/dto"
    "errors"
    "go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_UpdateRole (object *dto.Role)  error {

	result, err := dbConfig.DATABASE.Collection("Roles").UpdateOne(context.Background(), bson.M{"roleid": object.RoleId, "deleted":false}, bson.M{"$set": object})
	if err != nil {
		return err
	}
	if result.ModifiedCount < 1 && result.MatchedCount != 1 {
		return errors.New("Specified ID not found!")
	}

	return nil
}