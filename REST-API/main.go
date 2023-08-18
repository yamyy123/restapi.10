package main

import (
	"REST-API/config"
	"REST-API/constants"
	"REST-API/controllers"
	"REST-API/routes"
	"REST-API/services"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var ( //first we are declaring three global variables
	mongoClient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)


func initApp(mongoclient *mongo.Client) {
	ctx = context.TODO()
	transactionCollection := mongoclient.Database(constants.Database ).Collection("transaction")
	transactionServices := services. NewTransactionServiceInit(transactionCollection, ctx)
	transactionController := controllers.InitTransactionController(transactionServices)
	routes.TransactionRoute(server, transactionController)
}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDataBase()

	defer mongoclient.Disconnect(ctx)
	//this line is written to disconnect the server at last

	if err != nil {
		panic(err)
	}
	
	//initialise the app
	initApp(mongoclient)
	fmt.Println("Database connection initiated", constants.Port)

	log.Fatal(server.Run(constants.Port))
}
