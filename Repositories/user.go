package repositories

import (
	"context"
	"errors"

	"github.com/Abzaek/clean-arch/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoUserService implements the UserService interface using MongoDB.
type MongoUserService struct {
	collection *mongo.Collection
}

// NewMongoUserService creates a new MongoUserService with the given MongoDB collection.
func NewMongoUserService(db *mongo.Database, collectionName string) *MongoUserService {
	return &MongoUserService{
		collection: db.Collection(collectionName),
	}
}

//Finds a user

func (s *MongoUserService) Find(userId string) (*domain.User, error) {
	var user domain.User

	filter := bson.M{"_id": userId}

	singleResult := s.collection.FindOne(context.TODO(), filter)

	if singleResult.Decode(&user) != nil {
		return &user, errors.New("user doesn't exist")
	}

	return &user, nil
}

func (s *MongoUserService) FindMany() ([]*domain.User, error) {
	var result []*domain.User

	cur, err := s.collection.Find(context.TODO(), bson.D{{}})

	if err != nil {
		return []*domain.User{}, err
	}

	for cur.Next(context.TODO()) {
		var elem domain.User

		err := cur.Decode(&elem)

		if err != nil {
			return result, err
		}
		result = append(result, &elem)
	}

	return result, nil
}

// Save inserts a new user or updates an existing user in MongoDB.
func (s *MongoUserService) Save(user *domain.User) error {
	filter := bson.M{"_id": user.ID}

	singleResult := s.collection.FindOne(context.TODO(), filter)

	if singleResult.Err() == nil {
		return errors.New("user already exists")
	}

	manyResult, err := s.FindMany()

	if err != nil {
		return err
	}

	if len(manyResult) == 0 {
		user.Role = "admin"
	}

	update := bson.M{
		"$set": user,
	}

	_, err = s.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}

// Delete removes a user from MongoDB by ID.
func (s *MongoUserService) Delete(userId string) error {

	filter := bson.M{"_id": userId}
	_, err := s.collection.DeleteOne(context.Background(), filter)
	return err

}

// Update modifies an existing user in MongoDB.
func (s *MongoUserService) Update(user *domain.User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": user,
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	return err
}
