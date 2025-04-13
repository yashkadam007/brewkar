package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yashkadam007/brewkar/internal/controller"
)

// SetupRouter configures all routes for the application
func SetupRouter(
	authController *controller.AuthController,
	// Add more controllers as needed:
	// userController *controller.UserController,
	// beanController *controller.BeanController,
	// recipeController *controller.RecipeController,
	// brewLogController *controller.BrewLogController,
) *gin.Engine {
	// Set up Gin router
	router := gin.Default()

	// Register ping endpoint for health checks
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "pong",
		})
	})

	// Register API routes
	v1 := router.Group("/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authController.Register)
			auth.POST("/login", authController.Login)
			auth.POST("/refresh", authController.Refresh)
		}

		// Protected routes
		// api := v1.Group("")
		// TODO: Add auth middleware

		// User routes
		// users := api.Group("/users")
		// {
		// 	users.GET("/me", userController.GetProfile)
		// 	users.PUT("/me", userController.UpdateProfile)
		// }

		// // Bean routes
		// beans := api.Group("/beans")
		// {
		// 	beans.GET("", beanController.GetAll)
		// 	beans.POST("", beanController.Create)
		// 	beans.GET("/:id", beanController.GetByID)
		// 	beans.PUT("/:id", beanController.Update)
		// 	beans.DELETE("/:id", beanController.Delete)
		// 	beans.POST("/:id/images", beanController.UploadImages)
		// }

		// // Recipe routes
		// recipes := api.Group("/recipes")
		// {
		// 	recipes.GET("", recipeController.GetAll)
		// 	recipes.POST("", recipeController.Create)
		// 	recipes.GET("/:id", recipeController.GetByID)
		// 	recipes.PUT("/:id", recipeController.Update)
		// 	recipes.DELETE("/:id", recipeController.Delete)
		// 	recipes.POST("/:id/images", recipeController.UploadImages)
		// 	recipes.GET("/public", recipeController.GetPublic)
		// }

		// // Brew log routes
		// brewLogs := api.Group("/brew-logs")
		// {
		// 	brewLogs.GET("", brewLogController.GetAll)
		// 	brewLogs.POST("", brewLogController.Create)
		// 	brewLogs.GET("/:id", brewLogController.GetByID)
		// 	brewLogs.PUT("/:id", brewLogController.Update)
		// 	brewLogs.DELETE("/:id", brewLogController.Delete)
		// 	brewLogs.POST("/:id/images", brewLogController.UploadImages)
		// }

		// TODO: Add other routes (analytics, social, etc.)
	}

	return router
}
