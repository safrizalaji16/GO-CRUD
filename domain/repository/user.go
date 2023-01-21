package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/db"
	"github.com/safrizalaji16/learnGO/domain"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	ReadAllUsers(c *gin.Context, src *[]domain.User) (err error)
	CreateUser(c *gin.Context, src *domain.User) (err error)
	ReadDetailUser(c *gin.Context, src *domain.User, id int64) (err error)
	DeleteUser(c *gin.Context, src *domain.User, id int64) (err error)
	EditUser(c *gin.Context, src *domain.User, id int64) (err error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: db.GetConn(),
	}
}

func (r *userRepository) ReadAllUsers(c *gin.Context, src *[]domain.User) (err error) {
	if err = r.db.NewSelect().Model(src).Scan(c); err != nil {
		return err
	}

	return
}

func (r *userRepository) CreateUser(c *gin.Context, src *domain.User) (err error) {
	if _, err = r.db.NewInsert().Model(src).Exec(c); err != nil {
		return err
	}

	return
}

func (r *userRepository) ReadDetailUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if err = r.db.NewSelect().Model(src).Where("id = ?", &id).Scan(c); err != nil {
		return err
	}

	return
}

func (r *userRepository) DeleteUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if _, err = r.db.NewDelete().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}

func (r *userRepository) EditUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if _, err = r.db.NewUpdate().Model(src).Where("id = ?", &id).Exec(c); err != nil {
		return err
	}

	return
}
