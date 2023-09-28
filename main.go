package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/anushgowda/GoLang_Project/config"
	"github.com/anushgowda/GoLang_Project/controllers"
	"github.com/anushgowda/GoLang_Project/routes"
	"github.com/anushgowda/GoLang_Project/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	server = gin.Default()
	InitializeDatabase()
	InitializeProducts()
	InitializeUser()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)
	server.Run(":4000")
}

func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to Database %v", err)
	} else {
		fmt.Println("Connected to Database")
	}
}
func InitializeProducts() {
	productCollection := config.GetCollection(mongoClient, "anush_ekart", "products")
	productSvc := services.InitProductService(productCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)
}

func InitializeUser() {
	userCollection := config.GetCollection(mongoClient, "anush_ekart", "users")
	userSvc := services.InitUserService(userCollection)
	userController := controllers.InitAuthController(userSvc)
	routes.UserRoutes(server, *userController)
}