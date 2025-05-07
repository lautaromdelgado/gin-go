package processor

import "github.com/lautaromdelgado/gin-go/internal/User/dto"

type IUserProcessor interface {
	CreateUser(user *dto.UserRequest) error
}

type UserProcessor struct{}

func NewUserProcessor() IUserProcessor {
	return &UserProcessor{}
}
