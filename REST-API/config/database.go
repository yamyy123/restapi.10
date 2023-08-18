package config

import (
	"context"
	//"fmt"
	"log"
	"REST-API/constants"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDataBase() (*mongo.Client,error){
	ctx,_ :=context.WithTimeout(context.Background(),10* time.Second)
	//this line ensures whether the process takes place within 10 seconds

	mongoConnection :=options.Client().ApplyURI(constants.ConnectionString)
	//this helps us to get the mongoconnection from the server link which we have listed in the constants package by creating an instance of that mongodb
	//server

	mongoClient, err :=mongo.Connect(ctx, mongoConnection)
    //this helps us to connect to that particular server if it is not connected it will throw error

	if err!=nil{
		log.Fatal(err.Error())//if error came it will terminate the program
		return nil,err
	}
	if err:=mongoClient.Ping(ctx,readpref.Primary());err!=nil{//in this line we are seeing whether the particular server is active using the ping method
		return nil,err
	}
	return mongoClient,nil//here we are returning the server and no error by nil
}//this function will be called by 


func GetCollection(client *mongo.Client,dbName string,collectionName string) *mongo.Collection{
	// client,err :=ConnectDataBase()
	// if err!=nil{
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	collection :=client.Database(dbName).Collection(collectionName)
	return collection
}