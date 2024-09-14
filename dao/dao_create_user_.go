package dao

import (
    "context"
	"User-Mgt/dbConfig"
	"User-Mgt/dto"

)

func DB_CreateUser (object *dto.User) error {

	_, err := dbConfig.DATABASE.Collection("Users").InsertOne(context.Background(), object)
	if err != nil {
		return err
	}
	return nil
}