package services

import (
	"context"
	"encoding/json"
	"fmt"

	// "log"
	"mongodb/config"
	"mongodb/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

func ProductContent() *mongo.Collection{
    client,_:=config.ConnectDataBase()
    return config.GetCollection(client,"sample_analytics","transactions")
}

func FetchTransactions()([]*models.Transactions,error){
	ctx,_:= context.WithTimeout(context.Background(),100*time.Second)

	filter :=bson.M{"transaction_count":bson.D{{"$gt",85},{"$lt",90}}}
	//$gt-->gives greater than 85 and lower than 90

	//bson.M{}---stands for bson mapping
	//filter :=bson.D{}//this is used to get the requirements as our requirement
	//bson.D{}-->stands for bson document


	//options:=options.Find().SetLimit(100)//we can also do .sort
	//options := options.Find().SetSort(bson.D{{"transaction_count",1/*this will sort in asc order */}}).SetSkip(30)/*skip first 30*/.SetLimit(10)
	options := options.Find().SetSort(bson.D{{"transaction_count",-1/*this will sort in desc order */}}).SetSkip(30)/*skip first 30*/.SetLimit(10)
	result,err:=ProductContent().Find(ctx,filter,options)
    if err != nil {
        fmt.Println(err.Error())//it will display the textual representation of the error message.
        return nil, err
    } else {
		//build the array of products for the cursor that we have received
        var rest []*models.Transactions
        for result.Next(ctx) {
            post := &models.Transactions{}//creating a reference to store the posts as results
            err := result.Decode(post)//unmarshal(any to go structure)
            if err != nil {
                return nil, err
            }

            rest = append(rest, post)
        }
        if err := result.Err(); err != nil {
            return nil, err
        }
        return rest, nil
    }
}


func FetchAggregatedTransactions() {
	ctx,_:= context.WithTimeout(context.Background(),100*time.Second)
	matchStage :=bson.D{{"$match",bson.D{{"transaction_count",100}}}}
	groupStage:=bson.D{
		{"$group",bson.D{
			{"_id","$account_id"},
			{"total_count",bson.D{{"$sum","$transaction_count"}}},
		}}}
       result,err :=ProductContent().Aggregate(ctx,mongo.Pipeline{matchStage,groupStage})
	   if err ==nil{
		var showsWithinfo []bson.M
		if err = result.All(ctx, &showsWithinfo);err !=nil{
			panic(err)
		}
   formatted_data, err := json.MarshalIndent(showsWithinfo, "", " ")
        if err != nil {
            fmt.Println(err.Error())
        } else {
            fmt.Println(string(formatted_data))
        }
    }else{
		fmt.Println(err.Error())
	   }
}

 func UpdateTransaction(initialValue int,newValue int){
 	ctx,_ :=context.WithTimeout(context.Background(),100*time.Second)
 	filter := bson.D{{"account_id",initialValue}}
 	update := bson.D{{"$set",bson.D{{"account_id",newValue}}}}
 	result,err := ProductContent().UpdateOne(ctx,filter,update)
	if err!=nil{
		fmt.Println(err.Error())
		return nil,err
 	}
 	return result,nil
  }