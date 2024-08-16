package repositories

import (
	"context"
	"testing"
	"time"

	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskSuite struct {
	suite.Suite
	repo     usecases.TaskService
	dbClient *mongo.Client
}

func (suite *taskSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	dbClient, _ := mongo.Connect(context.TODO(), clientOptions)
	database := dbClient.Database("api")

	repo := NewMongoTaskService(database, "task")

	suite.repo = repo
	suite.dbClient = dbClient
}

func (suite *taskSuite) TearDownTest() {
	suite.dbClient.Database("api").Collection("task").DeleteMany(context.TODO(), bson.D{})
}

func (suite *taskSuite) TearDownSuite() {
	suite.dbClient.Disconnect(context.TODO())
}
func (suite *taskSuite) TestSave() {
	task := domain.Task{
		ID:          "12",
		Title:       "admin",
		Description: "hulsdf gjiskgsokg skddgjsoi",
		DueDate:     time.Now(),
		Status:      "done",
	}

	err := suite.repo.Save(&task)

	assert.NoError(suite.T(), err)
}

func (suite *taskSuite) TestGetById() {
	taskID := "12"

	task := domain.Task{
		ID:          taskID,
		Title:       "admin",
		Description: "hulsdf gjiskgsokg skddgjsoi",
		DueDate:     time.Now(),
		Status:      "done",
	}
	suite.repo.Save(&task)

	retrievedTask, err := suite.repo.GetById(taskID)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedTask)
	assert.Equal(suite.T(), taskID, retrievedTask.ID)
	assert.Equal(suite.T(), task.Title, retrievedTask.Title)
}

func (suite *taskSuite) TestUpdate() {
	taskID := "12"

	task := domain.Task{
		ID:          taskID,
		Title:       "admin",
		Description: "hulsdf gjiskgsokg skddgjsoi",
		DueDate:     time.Now(),
		Status:      "pending",
	}
	suite.repo.Save(&task)

	updatedTask := domain.Task{
		ID:          taskID,
		Title:       "super-admin",
		Description: "updated description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "done",
	}
	err := suite.repo.Update(&updatedTask)

	assert.NoError(suite.T(), err)

	retrievedTask, err := suite.repo.GetById(taskID)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedTask)
	assert.Equal(suite.T(), updatedTask.Title, retrievedTask.Title)
	assert.Equal(suite.T(), updatedTask.Description, retrievedTask.Description)
	assert.Equal(suite.T(), updatedTask.Status, retrievedTask.Status)
}

func (suite *taskSuite) TestDelete() {
	taskID := "12"

	task := domain.Task{
		ID:          taskID,
		Title:       "admin",
		Description: "hulsdf gjiskgsokg skddgjsoi",
		DueDate:     time.Now(),
		Status:      "done",
	}

	suite.repo.Save(&task)

	err := suite.repo.Delete(taskID)

	assert.NoError(suite.T(), err)

	retrievedTask, err := suite.repo.GetById(taskID)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), retrievedTask)
}

func (suite *taskSuite) TestGetAll() {
	task1 := domain.Task{
		ID:          "13",
		Title:       "task 1",
		Description: "description 1",
		DueDate:     time.Now(),
		Status:      "done",
	}
	task2 := domain.Task{
		ID:          "14",
		Title:       "task 2",
		Description: "description 2",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	suite.repo.Save(&task1)
	suite.repo.Save(&task2)

	tasks, err := suite.repo.GetAll()

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), tasks)
	assert.Len(suite.T(), tasks, 2)
}

func TestTaskSuite(t *testing.T) {
	suite.Run(t, new(taskSuite))
}
