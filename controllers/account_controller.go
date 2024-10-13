package controllers

import (
	"ecommerce-crm-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var accounts = []models.Account{
	{Id: 1, Name: "Clean Room", Email: "some@mail.ru", DarkMode: true},
}

func GetAccounts(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, accounts)
}
