package dao

import (
    "context"
	"User-Mgt/dbConfig"
	"User-Mgt/dto"

)

func DB_CreateRole (object *dto.Role) error {

	_, err := dbConfig.DATABASE.Collection("Roles").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}