package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/domain"
	"github.com/safrizalaji16/learnGO/services"
)

func ReadAllUsers(c *gin.Context) {
	var users []domain.User

	err := services.ReadAllUsers(c, &users)
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

// func CreateUser(c *gin.Context) {
// 	var src requests.CreateUserRequest

// 	if err := c.ShouldBindJSON(&src); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	var user User

// 	if err := copier.Copy(&user, &src); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	if _, err := db.NewInsert().Model(&user).Exec(c); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"data": src,
// 	})
// }

// func ReadDetailUser(c *gin.Context) {

// 	id := c.Param("id")
// 	var user User

// 	err := db.NewSelect().Model(&user).Where("id = ?", &id).Scan(c)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": user,
// 	})
// }

// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user User

// 	_, err := db.NewDelete().Model(&user).Where("id = ?", &id).Exec(c)
// 	if err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": `User has been deleted`,
// 	})
// }

// func EditUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var src CreateUserRequest

// 	if err := c.ShouldBindJSON(&src); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	var user User

// 	if err := copier.Copy(&user, &src); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	if _, err := db.NewUpdate().Model(&user).Where("id = ?", &id).Exec(c); err != nil {
// 		fmt.Println(err)
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "something error!",
// 		})

// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"data": src,
// 	})
// }
