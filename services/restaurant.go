 package services

// import (
//     "context"
//     "fmt"

//     // "log"
//     "mongodb/config"
//     "mongodb/models"
//     "time"

//     "go.mongodb.org/mongo-driver/bson"
//     "go.mongodb.org/mongo-driver/mongo"
//     "go.mongodb.org/mongo-driver/mongo/options"
//     // "go.mongodb.org/mongo-driver/mongo"
//     // "go.mongodb.org/mongo-driver/mongo/options"
// )
// func ProductCont() *mongo.Collection{
//     client,_:=config.ConnectDataBase()
//     return config.GetCollection(client,"sample_restaurants","restaurants")
    
// }
// func FindRes() ([]*models.Res, error) {
//     ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
//     filter := bson.D{}
//     result, err := ProductCont().Find(ctx, filter, options.Find().SetLimit(1))
//     if err != nil {
//         fmt.Println(err.Error())//it will display the textual representation of the error message.
//         return nil, err
//     } else {
//         var rest []*models.Res
//         for result.Next(ctx) {
//             post := &models.Res{}
//             err := result.Decode(post)
//             if err != nil {
//                 return nil, err
//             }

//             rest = append(rest, post)
//         }
//         if err := result.Err(); err != nil {
//             return nil, err
//         }
//         return rest, nil
//     }

// }

