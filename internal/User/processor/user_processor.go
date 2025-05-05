package processor

import "github.com/lautaromdelgado/gin-go/internal/User/domain"

type IUserProcessor interface {
	CreateUser(user *domain.User) error
}

type UserProcessor struct{}

func NewUserProcessor() IUserProcessor {
	return &UserProcessor{}
}
