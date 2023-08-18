package services

import (
	"REST-API/interfaces"
	"REST-API/models"
	"context"
	"errors"
	"time"

	//	"go.mongodb.org/mongo-driver/mango"
	"go.mongodb.org/mongo-driver/mongo"
)

type TransactionService struct{
	TransactionCollection *mongo.Collection
	ctx context.Context
}

func (p *TransactionService) CreateTransaction(user *models.Transaction) (*models.DBResponse, error) {
	
	
	user.Transaction_Date=time.Now()
	_, err := p.TransactionCollection.InsertOne(p.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}
	return nil, err
}
// func (p *TransactionService)GetTransaction(transaction *models.Transaction){
// 	panic("unimplemented")
// }
func NewTransactionServiceInit(collection *mongo.Collection,ctx context.Context) interfaces.ITransaction{
	return &TransactionService{ collection , ctx}
}


