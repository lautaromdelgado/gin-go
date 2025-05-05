package processor

import (
	"errors"

	"github.com/lautaromdelgado/gin-go/internal/User/domain"
)

func (u *UserProcessor) CreateUser(user *domain.User) error {
	if user.Nombre == "" {
		return errors.New("nombre is required")
	}
	if user.Apellido == "" {
		return errors.New("apellido is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	return nil
}
