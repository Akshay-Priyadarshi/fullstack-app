package abstractions

import "github.com/Akshay-Priyadarshi/fullstack-app/internal/app/dtos"

type IAuthService interface {
	Login(loginReqDto *dtos.AuthLoginReqData) (*dtos.AuthResData, error)
	Register(registerReqDto *dtos.AuthRegisterReqData) (*dtos.AuthResData, error)
}
