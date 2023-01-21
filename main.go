package main

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/controllers"
	"github.com/safrizalaji16/learnGO/db"
)

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	// db := db.GetConn()

	r.GET("users", controllers.ReadAllUsers)

	// r.POST("users", controllers.CreateUser)

	// r.GET("users/:id", controllers.ReadDetailUser)

	// r.DELETE("users/:id", controllers.DeleteUser)

	// r.PUT("users/:id", controllers.EditUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
