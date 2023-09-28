package services

import (
	"context"
	"fmt"
	// "errors"
	// "fmt"
	"strings"
	"time"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"github.com/anushgowda/GoLang_Project/utils"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	UserCollection *mongo.Collection
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

func (uc *UserService) Login(user *entities.Login) (string, error) {
	ctx := context.Background()
	query := bson.M{"Email": strings.ToLower(user.Email)}
	var loginResult *entities.Register
	err := uc.UserCollection.FindOne(ctx, query).Decode(&loginResult)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	//compare hashsed password with user entered password
	err2 := utils.VerifyPassword(loginResult.Password, user)
	if err != nil {
		return "wrong email or ", err2
	}
	return "login Successful", nil
}

func (uc *UserService) Logout() string{
	return "Logout Successful"
}