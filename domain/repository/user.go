package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/db"
	"github.com/safrizalaji16/learnGO/domain"
)

func ReadAllUsers(c *gin.Context, src *[]domain.User) (err error) {
	db := db.GetConn()

	if err = db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}
