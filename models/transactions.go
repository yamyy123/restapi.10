package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transactions struct {
	ID               primitive.ObjectID `json:"id" bson:"_id"`
	AccountID        int                `json:"accountid" bson:"account_id"`
	Transactioncount int                `json:"transactioncount" bson:"transaction_count"`
	Buckstartdate    primitive.DateTime `json:"bucketstartdate" bson:"bucket_start_date"`
	Buckenddate      primitive.DateTime `json:"bucketenddate" bson:"bucket_end_date"`
	Transactions     []Transaction      `json:"transactions" bson:"transactions"`
}

type Transaction struct {
	Date            primitive.DateTime `json:"date" bson:"date"`
	Amount          int                `json:"amount" bson:"amount"`
	Transactioncode string             `json:"transaction_code" bson:"transaction_code"`
	Symbol          string             `json:"symbol" bson:"symbol"`
	Price           string             `json:"price" bson:"price"`
	Total           string             `json:"total" bson:"total"`
}
