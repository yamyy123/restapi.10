package interfaces

import(
	"REST-API/models"
)

type ITransaction interface{
	CreateTransaction(transaction *models.Transaction)(*models.DBResponse, error)
	//GetTransaction(profile *models.Transaction)(*models.DBResponse, error)
}

