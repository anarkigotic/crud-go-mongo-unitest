package user_repository

import (
	"context"
	"time"

	"../../database"
	model "../../models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

func Create(user model.User) error {
	var err error
	_, err = collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}
	return nil

}
func Read() (model.Users, error) {
	var users model.Users
	filter := bson.D{}
	//returns us a cursor that we must run late
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var user model.User
		//we assign the cursor to the user variable
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)

	}

	return users, nil
}
func Update(user model.User, userId string) error {
	var err error
	//  convertimos ese user id string a un objeto id primitivp
	var oid, _ = primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": oid}
	update := bson.M{
		"$set": bson.M{
			"name":      user.Name,
			"email":     user.Email,
			"update_at": time.Now(),
		},
	}

	_, err = collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return err
	}

	return nil
}
func Delete(userId string) error {
	var err error
	var oid primitive.ObjectID

	//  convertimos ese user id string a un objeto id primitivp
	oid, _ = primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	var filter = bson.M{"_id": oid}

	_, err = collection.DeleteOne(ctx, filter)

	if err != nil {
		return err
	}

	return nil
}
