package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	/*	router := gin.New()
		router.Use(gin.Logger())
		router.Use(gin.Recovery())

		api := router.Group("/api")
		{
			user := api.Group("/user")
			user.GET("", init.UserCtrl.GetAllUserData)
				user.POST("", init.UserCtrl.AddUserData)
				user.GET("/:userID", init.UserCtrl.GetUserById)
				user.PUT("/:userID", init.UserCtrl.UpdateUserData)
				user.DELETE("/:userID", init.UserCtrl.DeleteUser)
		}*/

	router := gin.Default()

	// Account routes
	login := router.Group("/login")
	{
		login.GET("/", accountControler)
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	return router
}
