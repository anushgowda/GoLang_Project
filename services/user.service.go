package services

import (
	"context"
	// "errors"
	// "fmt"
	"strings"
	"time"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"github.com/anushgowda/GoLang_Project/utils"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	UserCollection *mongo.Collection
}

// Login implements interfaces.IUser.
func (*UserService) Login(user *entities.Login) (string, error) {
	panic("unimplemented")
}

// Logout implements interfaces.IUser.
func (*UserService) Logout(error) {
	panic("unimplemented")
}

func InitUserService(collection *mongo.Collection) interfaces.IUser {
	return &UserService{UserCollection: collection}
}

func (uc *UserService) Register(user *entities.Register) (string, error) {
	// ctx := context.Background()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)

	if user.Password != user.PasswordConfirm{
		return "password matched",nil

	}

	if hashPassword := utils.EncryptPassword(user); hashPassword != nil {
	  user.Password = string(hashPassword)
	  user.PasswordConfirm = string(hashPassword)
	} else {
	    return "Error in Password Encryption", nil
	}
	_, err :=  uc.UserCollection.InsertOne(context.Background(), user)
	if err != nil {

        return "", err

    } else {

        return "User registered Successsfully", nil

    }
}

// func (uc *UserService) Login(user *entities.Login) (*entities.LoginResponse, error) {
// 	ctx := context.Background()
// 	query := bson.M{"Email": strings.ToLower(user.Email)}
// 	var loginResult *entities.User
// 	err := uc.UserCollection.FindOne(ctx, query).Decode(&loginResult)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return nil, err
// 	}
// 	//compare hashsed password with user entered password
// 	err2 := utils.VerifyPassword(loginResult.Password, user.Password)
// 	if err != nil {
// 		return nil, err2
// 	}
// 	return &entities.LoginResponse{Token: token, RefreshToken: refreshToken}, nil
// }
