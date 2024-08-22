package service

import "github.com/carmenphan/go-ecommerce-backend-api/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}
func (us *UserService) GetUserByIdService() string {
	return us.UserRepo.GetIoUser()

}
