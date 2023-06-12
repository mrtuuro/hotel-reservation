package db

import (
	"context"
	"fmt"
	"github.com/mrtuuro/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userColl = "users"
)

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserByID(context.Context, string) (*types.User, error)
	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUserByID(context.Context, string) error
	UpdateUserByID(ctx context.Context, filter bson.M, update types.UpdateUserParams) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func (m *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("Database dropping user collection")
	return m.coll.Drop(ctx)
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (m *MongoUserStore) UpdateUserByID(ctx context.Context, filter bson.M, values types.UpdateUserParams) error {
	update := bson.D{
		{
			"$set", values.ToBSON(),
		},
	}
	_, err := m.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) DeleteUserByID(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// TODO if no user is deleted we may log it.
	_, err = m.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (m *MongoUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	var user types.User
	if err := m.coll.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user types.User
	if err := m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cursor, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.coll.InsertOne(ctx, user, nil)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil

}
