package routes

import (
	"github.com/gin-gonic/gin"
	"go-crud-v1/controllers"
)

func SetupRoutes() {
	router := gin.Default()

	api := router.Group("/api/v1/")
	{
		api.POST("create-post", controllers.PostCreate)
		api.PUT("update-post", controllers.PostUpdate)

		api.GET("posts", controllers.FindAll)
		api.GET("posts/:id", controllers.FindById)

		api.DELETE("post/:id", controllers.DeleteById)
	}

	router.Run()
}
