package routes

import (
	accountController "ecommerce-crm-server/controllers"
	"github.com/gin-gonic/gin"
)

func accountRoutes(superRoute *gin.RouterGroup) {
	accountRouter := superRoute.Group("/accounts")
	{

		accountRouter.GET("/", accountController.GetAccounts)
	}
}
