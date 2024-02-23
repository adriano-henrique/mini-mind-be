package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func BuildRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	versionRouter := r.Group("/api/v1")
	nuggetsGroupRouter(versionRouter)
	folderGroupRouter(versionRouter)

	return r
}
