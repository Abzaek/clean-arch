package repositories

import (
	"context"
	"errors"

	"github.com/Abzaek/clean-arch/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTaskService struct {
	collection *mongo.Collection
}

// NewMongoTaskService creates a new MongoTaskService with the given MongoDB collection.
func NewMongoTaskService(db *mongo.Database, collectionName string) *MongoTaskService {
	return &MongoTaskService{
		collection: db.Collection(collectionName),
	}
}

// Save inserts a new task or updates an existing task in MongoDB.
func (s *MongoTaskService) Save(task *domain.Task) error {
	filter := bson.M{"_id": task.ID}

	singleResult := s.collection.FindOne(context.TODO(), filter)

	if singleResult.Err() == nil {
		return errors.New("task already exists")
	}

	update := bson.M{
		"$set": task,
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true))
	return err
}

// Update modifies an existing task in MongoDB.
func (s *MongoTaskService) Update(task *domain.Task) error {
	filter := bson.M{"_id": task.ID}
	update := bson.M{
		"$set": task,
	}
	_, err := s.collection.UpdateOne(context.Background(), filter, update)
	return err
}

// Delete removes a task from MongoDB by ID.
func (s *MongoTaskService) Delete(taskId string) error {
	filter := bson.M{"_id": taskId}
	_, err := s.collection.DeleteOne(context.Background(), filter)
	return err
}

// GetById retrieves a task from MongoDB by ID.
func (s *MongoTaskService) GetById(taskId string) (*domain.Task, error) {
	filter := bson.M{"_id": taskId}
	var task domain.Task
	err := s.collection.FindOne(context.Background(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// GetAll retrieves all tasks from MongoDB.
func (s *MongoTaskService) GetAll() ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor, err := s.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
