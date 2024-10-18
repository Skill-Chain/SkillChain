package routes

import "github.com/gin-gonic/gin"

func addAuthGroup(rg *gin.RouterGroup) {
	auth := rg.Group("auth")

	// можно вынести запрос и для каждого запроса создать переменную, потом просто
	// перенести переменную метода в AddAuthGroup
	auth.GET("/", func(c *gin.Context) {

	})

	auth.POST("/login", func(c *gin.Context) {

	})

	auth.POST("/logout", func(c *gin.Context) {

	})

	auth.POST("/register", func(c *gin.Context) {

	})

	auth.POST("/req-forgot-password", func(c *gin.Context) {

	})

	auth.POST("/reset-password", func(context *gin.Context) {

	})
}
