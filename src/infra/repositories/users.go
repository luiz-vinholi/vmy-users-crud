package repositories

import (
	"context"
	"time"
	"vmytest/src/infra/databases"
	"vmytest/src/infra/models"

	"github.com/chidiwilliams/flatbson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsersRepository struct {
	client *mongo.Collection
}

func NewUsersRepository() *UsersRepository {
	client := databases.Mongo.Collection("users")
	return &UsersRepository{client}
}

func (ur *UsersRepository) GetUsers() ([]models.User, error) {
	ctx, cancel := getContext()
	defer cancel()

	cursor, err := ur.client.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (ur *UsersRepository) GetUser(id string) (*models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objId}
	user, err := ur.getUser(filter)
	return user, err
}

func (ur *UsersRepository) GetUserByEmail(email string) (*models.User, error) {
	filter := bson.M{"email": email}
	user, err := ur.getUser(filter)
	return user, err
}

func (ur *UsersRepository) CreateUser(data models.User) (id string, err error) {
	ctx, cancel := getContext()
	defer cancel()

	data.CreatedDate = time.Now()
	result, err := ur.client.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (ur *UsersRepository) UpdateUser(id string, data models.User) (err error) {
	ctx, cancel := getContext()
	defer cancel()

	data.UpdatedDate = time.Now()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	flattedData, err := flatbson.Flatten(data)
	if err != nil {
		return
	}

	_, err = ur.client.UpdateByID(ctx, objId, bson.M{"$set": flattedData})
	if err != nil {
		return
	}
	return
}

func (ur *UsersRepository) DeleteUser(id string) (err error) {
	ctx, cancel := getContext()
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return
	}
	filter := bson.M{"_id": objId}

	_, err = ur.client.DeleteOne(ctx, filter)
	if err != nil {
		return
	}
	return
}

func (ur *UsersRepository) getUser(filter bson.M) (*models.User, error) {
	ctx, cancel := getContext()
	defer cancel()

	var result models.User
	if err := ur.client.FindOne(ctx, filter).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &result, nil
}

func getContext() (ctx context.Context, cancel func()) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}
