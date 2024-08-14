package routers

import (
	"github.com/Abzaek/clean-arch/Infrastructure"
	"github.com/Abzaek/clean-arch/delivery/controllers"
	"github.com/gin-gonic/gin"
)

func StartApp(tCtrl *controllers.TaskController, uCtrl *controllers.UserController, router *gin.Engine, authM *Infrastructure.AuthMiddleware) {

	router.POST("/register", uCtrl.RegisterUser)
	router.POST("/login", uCtrl.LoginUser)

	middle := router.Group("")

	middle.Use(authM.ValidateToken())

	middle.GET("/tasks", tCtrl.GetTasks)
	middle.GET("/tasks/:id", tCtrl.GetTaskById)
	middle.POST("/tasks", tCtrl.CreateTask)
	middle.DELETE("/tasks/:id", tCtrl.DeleteTask)
	middle.PUT("/tasks", tCtrl.UpdateTask)

	middle.POST("/promote", uCtrl.PromoteUser)
}
