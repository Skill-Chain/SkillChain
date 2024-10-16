package routes

import "github.com/gin-gonic/gin"

var Router = gin.Default()

// GetRoutes - Штука, которая отвечает за группировку url, endpoint
func GetRoutes() {
	groupUser := Router.Group("/public/user")
	groupEmployer := Router.Group("/public/employer")
	groupAuth := Router.Group("/public/auth")

	addAuthGroup(groupAuth)
	addUserRoutes(groupUser)
	addEmployerRoutes(groupEmployer)

}
