package repositories

import (
	"context"
	"testing"

	"github.com/Abzaek/clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userSuite struct {
	suite.Suite
	repo     *MongoUserService
	dbClient *mongo.Client
}

func (suite *userSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	dbClient, _ := mongo.Connect(context.TODO(), clientOptions)

	repo := NewMongoUserService(dbClient.Database("api"), "user")

	suite.repo = repo
	suite.dbClient = dbClient
}

func (suite *userSuite) TearDownTest() {
	// Clear the collection after each test
	suite.repo.collection.DeleteMany(context.TODO(), bson.D{})
}

func (suite *userSuite) TearDownSuite() {
	// Close the MongoDB client connection after all tests are done
	if err := suite.dbClient.Disconnect(context.TODO()); err != nil {
		suite.T().Fatal("Failed to close the database connection:", err)
	}
}

func (suite *userSuite) TestSave() {
	user := domain.User{
		ID:       "1",
		Role:     "admin",
		Password: "1234",
	}

	err := suite.repo.Save(&user)

	assert.NoError(suite.T(), err)
}

func (suite *userSuite) TestDelete() {
	userId := "1"
	user := domain.User{
		ID:       userId,
		Role:     "admin",
		Password: "1234",
	}

	// Save the user
	suite.repo.Save(&user)

	// Ensure the user was saved correctly
	retrievedUser, err := suite.repo.Find(userId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.ID, retrievedUser.ID)

	// Delete the user
	err = suite.repo.Delete(userId)
	assert.NoError(suite.T(), err)

	// Ensure the user was deleted
	_, err = suite.repo.Find(userId)
	assert.Error(suite.T(), err)
}

func (suite *userSuite) TestUpdate() {
	userId := "1"
	user := domain.User{
		ID:       userId,
		Role:     "admin",
		Password: "1234",
	}

	// Save the initial user
	suite.repo.Save(&user)

	// Update the user's role
	updatedUser := domain.User{
		ID:       userId,
		Role:     "super-admin",
		Password: "1234",
	}
	err := suite.repo.Update(&updatedUser)

	assert.NoError(suite.T(), err)

	// Retrieve the updated user and verify the changes
	retrievedUser, err := suite.repo.Find(userId)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), updatedUser.Role, retrievedUser.Role)
	assert.Equal(suite.T(), updatedUser.Password, retrievedUser.Password)
}

func (suite *userSuite) TestFind() {
	userId := "1"
	user := domain.User{
		ID:       userId,
		Role:     "admin",
		Password: "1234",
	}

	// Save the user
	suite.repo.Save(&user)

	// Attempt to find the user by ID
	retrievedUser, err := suite.repo.Find(userId)

	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), retrievedUser)
	assert.Equal(suite.T(), user.ID, retrievedUser.ID)
	assert.Equal(suite.T(), user.Role, retrievedUser.Role)
}

func TestUserSuite(t *testing.T) {
	suite.Run(t, new(userSuite))
}
