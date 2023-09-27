package interfaces

import "github.com/anushgowda/GoLang_Project/entities"

type IUser interface {
	Register(user *entities.User) (string, error)
	Login(user *entities.Login) (string, error)
	Logout(error)
}