package routes

import "github.com/gin-gonic/gin"

func addEmployerRoutes(rg *gin.RouterGroup) {
	user := rg.Group("/user")

	user.GET("/profile/:id", func(c *gin.Context) {

	})
}
