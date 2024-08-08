package routes

import (
    "github.com/wonderstone/harpoon/controllers"
    "github.com/wonderstone/harpoon/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    protected := router.Group("/protected")
    protected.Use(middleware.AuthMiddleware())
    protected.GET("/dashboard", controllers.ProtectedEndpoint)
}