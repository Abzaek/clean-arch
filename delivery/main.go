package main

import (
	"context"
	"log"

	"github.com/Abzaek/clean-arch/Infrastructure"
	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/Abzaek/clean-arch/delivery/controllers"
	"github.com/Abzaek/clean-arch/delivery/routers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	router := gin.Default()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
	}

	database := dbClient.Database("api")

	userCollection := Infrastructure.NewMongoUserService(database, "user")
	taskCollection := Infrastructure.NewMongoTaskService(database, "task")

	userUseCase := usecases.NewUserUseCase(userCollection)
	taskUseCase := usecases.NewTaskUseCase(taskCollection)
	passM := &Infrastructure.PasswordServiceBcrypt{}
	auth := &Infrastructure.JwtService{
		JwtKey:  []byte("abzaeko"),
		Service: userCollection,
	}

	userC := &controllers.UserController{
		UUC:        userUseCase,
		UserAuth:   auth,
		PassManage: passM,
	}

	taskC := &controllers.TaskController{
		TUC: taskUseCase,
	}

	authM := &Infrastructure.AuthMiddleware{
		Auth: auth,
	}

	routers.StartApp(taskC, userC, router, authM)

	router.Run("localhost:3000")
}
