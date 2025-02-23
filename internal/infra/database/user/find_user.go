package user

import (
	"context"
	"errors"
	"github.com/liberopassadorneto/auction/configuration/logger"
	"github.com/liberopassadorneto/auction/internal/entity/user_entity"
	"github.com/liberopassadorneto/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (
	*user_entity.User,
	*internal_error.InternalError,
) {
	var user UserMongo
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error("user not found in mongodb", err)
			return nil, internal_error.NewNotFoundError("user not found in mongodb")
		}
		logger.Error("error finding user in mongodb", err)
		return nil, internal_error.NewInternalServerError("error finding user in mongodb")
	}

	return &user_entity.User{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
