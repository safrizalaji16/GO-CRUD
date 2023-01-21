package main

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/controllers"
	"github.com/safrizalaji16/learnGO/db"
	"github.com/safrizalaji16/learnGO/domain/repository"
	"github.com/safrizalaji16/learnGO/services"
)

func main() {
	r := gin.Default()

	dbInstance := db.InitDB()
	defer dbInstance.Close()

	rp := repository.InitRepositoryInstance()
	userSrv := services.NewUserService(rp.User)
	userCtl := controllers.NewUserController(userSrv)

	r.GET("users", userCtl.ReadAllUsers)

	r.POST("users", userCtl.CreateUser)

	r.GET("users/:id", userCtl.ReadDetailUser)

	r.DELETE("users/:id", userCtl.DeleteUser)

	r.PUT("users/:id", userCtl.EditUser)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
