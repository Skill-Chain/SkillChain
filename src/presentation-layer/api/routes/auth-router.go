package routes

import (
	"APImod/src/presentation-layer/api/controllers"
	"github.com/gin-gonic/gin"
)

func addAuthGroup(rg *gin.RouterGroup) {
	auth := rg.Group("auth")
	controller := &controllers.ControllerAuthResult{}
	// можно вынести запрос и для каждого запроса создать переменную, потом просто
	// перенести переменную метода в AddAuthGroup
	auth.POST("/register", controller.Register())
	auth.POST("/login", controller.Login())

	auth.GET("/", func(c *gin.Context) {
	})

	auth.POST("/logout", func(c *gin.Context) {
	})

	auth.POST("/req-forgot-password", func(c *gin.Context) {
	})

	auth.POST("/reset-password", func(context *gin.Context) {
	})
}
