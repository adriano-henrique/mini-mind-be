package routes

import (
	"github.com/adriano-henrique/mini-mind-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func folderGroupRouter(baseRouter *gin.RouterGroup) {
	folder := baseRouter.Group("/folder")

	folder.POST("/create", handler.CreateFolder)
	folder.GET("/all", handler.GetAllFolders)
}
