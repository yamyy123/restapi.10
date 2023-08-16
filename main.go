package main

import (
	//"context"
	"fmt"
	//"mongodb/constants"
	//"mongodb/models"
	"mongodb/services"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

var(
	mongoClient *mongo.Client
)
func main(){
	// ctx := context.TODO()

	// //connect to MongoDB
	// mongoconn :=options.Client().ApplyURI(constants.ConnectionString)
	// mongoclient,err :=mongo.Connect(ctx,mongoconn)
	// if err !=nil{
	// 	panic(err)
	// }

	// if err :=mongoclient.Ping(ctx, readpref.Primary());err !=nil{
	// 	panic(err)
	// }

	//client:= config.ConnectDataBase()
	//collection :=config.Getcollection(client,"sample_training","zips")
	fmt.Println("MongoDB successfully connected....")
//	product:=models.Product{ID: primitive.NewObjectID(),Name:}
    // products :=[]interface{}{
	// 	models.Product{ID: primitive.NewObjectID(),Name:"Oneplus",Price:1000000,Description: "Budget Phone"},
	// 	models.Product{ID: primitive.NewObjectID(),Name:"Vivo",Price:1000000,Description: "china Phone"},
	// }
	// services.InsertproductList(products)

	products,_ :=services.FindProducts()
	for _,product :=range products{
		fmt.Println(product)
	}
}