package routes

import (
	"example/controllers"
	"example/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// Tạo group protected
	protected := router.Group("/")
	protected.Use(middlewares.JWTAuth())
	{
		// Các endpoint cần bảo vệ
		protected.POST("/employee", controllers.CreateEmployee())
		protected.PUT("/employee/:id", controllers.UpdateEmployee())
		protected.GET("/employee", controllers.GetListEmployee())
		protected.GET("/employee/:id", controllers.GetDetailEmployee())
		protected.DELETE("/employee/:id", controllers.DeleteEmployee())

		protected.POST("/topic", controllers.CreateTopic())
		protected.PUT("/topic/:id", controllers.UpdateTopic())
		protected.GET("/topic", controllers.GetListTopic())
		protected.GET("/topic/:id", controllers.GetDetailTopic())
		protected.DELETE("topic/:id", controllers.DeleteEmployee())

		protected.POST("/statement", controllers.CreateStatement())
		protected.PUT("/statement/:id", controllers.UpdateStatement())
		protected.GET("/statement", controllers.GetListStatement())
		protected.GET("/statement/:id", controllers.GetDetailStatement())
		protected.DELETE("statement/:id", controllers.DeleteStatement())
		protected.GET("/statement/pending", controllers.GetListStatementPending())

		protected.GET("dashboard/statistic", controllers.StatisticDataInYear())

		protected.POST("upload", controllers.UploadFile())

	}

	// Các endpoint không cần bảo vệ
	router.POST("/login", controllers.LoginAdmin())
}
