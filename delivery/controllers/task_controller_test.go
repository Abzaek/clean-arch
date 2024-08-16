package controllers

import (
	"github.com/Abzaek/clean-arch/Usecases/mocks"
	"github.com/stretchr/testify/suite"
)

type tControllerSuite struct {
	suite.Suite
	usecase     *mocks.Usecases
	tController *TaskController
}

func (suite *tControllerSuite) SetupTest() {
	suite.usecase = &mocks.Usecases{}
	suite.tController = &TaskController{TUC: suite.usecase}

}
