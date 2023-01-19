package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/safrizalaji16/learnGO/db"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel
	ID       int64  `bun:"id,pk,autoincrement" json:"id"`
	Email    string `bun:"email" json:"email"`
	Password string `bun:"password" json:"password"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	db := db.GetConn()

	r.GET("/users", func(c *gin.Context) {
		users := []User{}
		err := db.NewSelect().Model(&users).Scan(c)
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
	})

	r.POST("/users", func(c *gin.Context) {
		var src CreateUserRequest

		if err := c.ShouldBindJSON(&src); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		var user User

		if err := copier.Copy(&user, &src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		if _, err := db.NewInsert().Model(&user).Exec(c); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": src,
		})
	})

	r.GET("/users/:id", func(c *gin.Context) {

		id := c.Param("id")
		var user User

		err := db.NewSelect().Model(&user).Where("id = ?", &id).Scan(c)
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
	})

	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user User

		_, err := db.NewDelete().Model(&user).Where("id = ?", &id).Exec(c)
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
	})

	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var src CreateUserRequest

		if err := c.ShouldBindJSON(&src); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		var user User

		if err := copier.Copy(&user, &src); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		if _, err := db.NewUpdate().Model(&user).Where("id = ?", &id).Exec(c); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something error!",
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"data": src,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
