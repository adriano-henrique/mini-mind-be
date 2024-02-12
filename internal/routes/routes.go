package routes

import "github.com/gin-gonic/gin"

func BuildRoutes() *gin.Engine {
	r := gin.Default()

	versionRouter := r.Group("/api/v1")
	nuggetsGroupRouter(versionRouter)

	return r
}
