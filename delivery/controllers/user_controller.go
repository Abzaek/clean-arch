package controllers

import (
	"net/http"

	usecases "github.com/Abzaek/clean-arch/Usecases"
	"github.com/Abzaek/clean-arch/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UUC        *usecases.UserUsecase
	UserAuth   usecases.AuthService
	PassManage usecases.PasswordService
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := (uc.PassManage).GenerateHash(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = hashedPassword
	user.Role = "user"

	err = uc.UUC.SaveUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := (uc.UserAuth).GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) LoginUser(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	existingUser, err := uc.UUC.FindUser(user.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	equals := (uc.PassManage).ComparePassword(existingUser.Password, user.Password)

	if !equals {
		c.JSON(http.StatusNotFound, "incorrect password")
		return
	}

	token, err := (uc.UserAuth).GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (uc *UserController) PromoteUser(c *gin.Context) {
	id := c.Param("id")

	existingUser, err := uc.UUC.FindUser(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	existingUser.Role = "admin"

	err = uc.UUC.UpdateUser(existingUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "updated successfully")
}
