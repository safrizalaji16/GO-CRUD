package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/domain"
	"github.com/safrizalaji16/learnGO/domain/repository"
)

func ReadAllUsers(c *gin.Context, src *[]domain.User) (err error) {
	if err = repository.ReadAllUsers(c, src); err != nil {
		return err
	}

	return
}

// func CreateUser(c *gin.Context, src)  {

// }
