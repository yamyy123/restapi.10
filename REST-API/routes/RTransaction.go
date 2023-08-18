package routes

import (
	"REST-API/controllers"
	//"rest-api/controllers"

	"github.com/gin-gonic/gin"
)

func TransactionRoute(router *gin.Engine,controller controllers.TransactionController){
	router.POST("/api/transaction/create", controller.CreateTransaction)
	//router.POST("/api/profile/edit", controller.GetTransaction)
}