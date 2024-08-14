package controllers

import (
	"github.com/gin-gonic/gin"
)

type Controllers interface {
	UpdateTask(c *gin.Context)
	CreateTask(c *gin.Context)
	DeleteTask(c *gin.Context)
	GetTasks(c *gin.Context)
	GetTaskById(c *gin.Context)

	LoginUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	PromoteUser(c *gin.Context)
	RegisterUser(c *gin.Context)
}
