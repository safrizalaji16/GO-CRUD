package services

import (
	"github.com/gin-gonic/gin"
	"github.com/safrizalaji16/learnGO/domain"
	"github.com/safrizalaji16/learnGO/domain/repository"
)

type UserService interface {
	ReadAllUsers(c *gin.Context, src *[]domain.User) (err error)
	CreateUser(c *gin.Context, src *domain.User) (err error)
	ReadDetailUser(c *gin.Context, src *domain.User, id int64) (err error)
	DeleteUser(c *gin.Context, src *domain.User, id int64) (err error)
	EditUser(c *gin.Context, src *domain.User, id int64) (err error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (srv *userService) ReadAllUsers(c *gin.Context, src *[]domain.User) (err error) {
	if err = srv.userRepo.ReadAllUsers(c, src); err != nil {
		return err
	}

	return
}

func (srv *userService) CreateUser(c *gin.Context, src *domain.User) (err error) {
	if err = srv.userRepo.CreateUser(c, src); err != nil {
		return err
	}

	return
}

func (srv *userService) ReadDetailUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if err = srv.userRepo.ReadDetailUser(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *userService) DeleteUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if err = srv.userRepo.DeleteUser(c, src, id); err != nil {
		return err
	}

	return
}

func (srv *userService) EditUser(c *gin.Context, src *domain.User, id int64) (err error) {
	if err = srv.userRepo.EditUser(c, src, id); err != nil {
		return err
	}

	return
}
