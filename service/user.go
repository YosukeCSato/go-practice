package service

import (
	"fmt"

	"../repository"
)

type UserService interface {
	GetUser(uint64) uint64
}

type userService struct {
	ur repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
	return userService{ur}
}

func (us userService) GetUser(uID uint64) uint64 {
	fmt.Println("---------------------")
	fmt.Println("---------------------")
	return us.ur.GetUser(uID)
}
