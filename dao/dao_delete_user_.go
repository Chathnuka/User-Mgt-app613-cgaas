package dao

import (
	"User-Mgt/dbConfig"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func DB_DeleteUser (userId string)  error {
  
        result, err := dbConfig.DATABASE.Collection("Users").UpdateOne(context.Background(), bson.M{"userid": userId}, bson.M{"$set": bson.M{"deleted": true}})
        if err != nil {
            return err
        }
        if result.ModifiedCount < 1 {
            return errors.New("Specified Id not found!")
        }
        return nil
  
}

