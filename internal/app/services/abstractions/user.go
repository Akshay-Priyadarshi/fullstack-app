package abstractions

import "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"

type IUserService interface {
	UpdatePassword(userPasswordUpdateDto *dtos.UserPasswordUpdateReqData) (*dtos.UserResData, error)

	GetById(id string) (*dtos.UserResData, error)
}
