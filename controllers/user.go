package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/safrizalaji16/learnGO/domain"
	"github.com/safrizalaji16/learnGO/dto/requests"
	"github.com/safrizalaji16/learnGO/services"
)

type UserController interface {
	ReadAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	ReadDetailUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	EditUser(c *gin.Context)
}
type userController struct {
	userSrv services.UserService
}

func NewUserController(userSrv services.UserService) UserController {
	return &userController{
		userSrv: userSrv,
	}
}

func (ctl *userController) ReadAllUsers(c *gin.Context) {
	var users []domain.User

	err := ctl.userSrv.ReadAllUsers(c, &users)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (ctl *userController) CreateUser(c *gin.Context) {
	var src requests.CreateUserRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var user domain.User

	if err := copier.Copy(&user, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	err := ctl.userSrv.CreateUser(c, &user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": src,
	})
}

func (ctl *userController) ReadDetailUser(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var user domain.User

	err = ctl.userSrv.ReadDetailUser(c, &user, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (ctl *userController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var user domain.User

	err = ctl.userSrv.DeleteUser(c, &user, id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": `User has been deleted`,
	})
}

func (ctl *userController) EditUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}
	var src requests.CreateUserRequest

	if err := c.ShouldBindJSON(&src); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	var user domain.User

	if err := copier.Copy(&user, &src); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	if err = ctl.userSrv.EditUser(c, &user, id); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "something error!",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": src,
	})
}
