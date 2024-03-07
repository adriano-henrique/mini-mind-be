package routes

import (
	"github.com/adriano-henrique/mini-mind-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func nuggetsGroupRouter(baseRouter *gin.RouterGroup) {
	nuggets := baseRouter.Group("/nuggets")

	nuggets.GET("/all", handler.GetAllNuggets)
	nuggets.GET("/get/:id", handler.GetNuggetById)
	nuggets.POST("/create", handler.CreateNugget)
	nuggets.PUT("/update", handler.UpdateNugget)
	nuggets.PATCH("/update", handler.UpdateNugget)
	nuggets.DELETE("/delete/:id", handler.DeleteNugget)
	nuggets.GET("/get_by_folder/:folder_id", handler.GetNuggetByFolderID)
	nuggets.GET("/get_by_key/:key", handler.GetNuggetByKey)
}
