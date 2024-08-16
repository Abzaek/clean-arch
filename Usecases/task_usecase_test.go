package usecases

import (
	"testing"
	"time"

	"github.com/Abzaek/clean-arch/Usecases/mocks"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type taskUseCaseSuite struct {
	suite.Suite
	tUsecase *TaskUseCase
	repo     *mocks.TaskService
}

func (suite *taskUseCaseSuite) SetupTest() {
	suite.repo = &mocks.TaskService{}
	suite.tUsecase = NewTaskUseCase(suite.repo)
}

func (suite *taskUseCaseSuite) TestGetAllTasks() {
	// Mock data
	tasks := []*domain.Task{
		{
			ID:          "1",
			Title:       "Task 1",
			Description: "Description 1",
			DueDate:     time.Now(),
			Status:      "pending",
		},
		{
			ID:          "2",
			Title:       "Task 2",
			Description: "Description 2",
			DueDate:     time.Now(),
			Status:      "completed",
		},
	}

	// Set up expected behavior
	suite.repo.On("GetAll").Return(tasks, nil)

	// Execute the use case
	result, err := suite.tUsecase.GetAllTasks()

	// Verify expectations
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), tasks, result)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestGetTaskById() {
	// Mock data
	task := &domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	// Set up expected behavior
	suite.repo.On("GetById", "1").Return(task, nil)

	// Execute the use case
	result, err := suite.tUsecase.GetTaskById("1")

	// Verify expectations
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), task, result)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestDeleteTask() {
	// Set up expected behavior
	suite.repo.On("Delete", "1").Return(nil)

	// Execute the use case
	err := suite.tUsecase.DeleteTask("1")

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestSaveTask() {
	// Mock data
	task := &domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	// Set up expected behavior
	suite.repo.On("Save", task).Return(nil)

	// Execute the use case
	err := suite.tUsecase.SaveTask(task)

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func (suite *taskUseCaseSuite) TestUpdateTask() {
	// Mock data
	task := &domain.Task{
		ID:          "1",
		Title:       "Task 1",
		Description: "Description 1",
		DueDate:     time.Now(),
		Status:      "pending",
	}

	// Set up expected behavior
	suite.repo.On("Update", task).Return(nil)

	// Execute the use case
	err := suite.tUsecase.UpdateTask(task)

	// Verify expectations
	assert.NoError(suite.T(), err)

	// Ensure that the expectations were met
	suite.repo.AssertExpectations(suite.T())
}

func TestTaskUseCaseSuite(t *testing.T) {
	suite.Run(t, new(taskUseCaseSuite))
}
