package controllers

import (
	"log"
	"net/http"

	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TUC usecases.Usecases
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	claims, exists := c.Get("Claims")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claim doesnt exist"})
		return
	}

	if claim, ok := claims.(*jwt.MapClaims); ok {
		if (*claim)["role"] != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			return
		}

	} else {
		log.Fatal("here lese")
		c.JSON(http.StatusBadRequest, "invalid claims bad request")
		return
	}

	var task domain.Task

	if err := c.BindJSON(&task); err != nil {
		log.Fatal("here task")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TUC.SaveTask(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		// log.Fatal("here err")
		return
	}

	c.JSON(http.StatusOK, "successfully created")
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	claims, exists := c.Get("Claims")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claim doesnt exist"})
		return
	}

	if claim, ok := claims.(*jwt.MapClaims); ok {
		if (*claim)["role"] != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			return
		}
	}

	var task domain.Task

	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TUC.UpdateTask(&task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "successfully Updated")
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	claims, exists := c.Get("Claims")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "claim doesnt exist"})
		return
	}

	if claim, ok := claims.(*jwt.MapClaims); ok {
		if (*claim)["role"] != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
			return
		}
	}

	id := c.Param("id")

	err := tc.TUC.DeleteTask(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

}

func (tc *TaskController) GetTasks(c *gin.Context) {
	tasks, err := tc.TUC.GetAllTasks()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")

	task, err := tc.TUC.GetTaskById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, task)
}
