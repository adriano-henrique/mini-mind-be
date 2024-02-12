package routes

import (
	"github.com/adriano-henrique/mini-mind-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func nuggetsGroupRouter(baseRouter *gin.RouterGroup) {
	nuggets := baseRouter.Group("/nuggets")

	nuggets.GET("/all", handler.GetAllNuggets)
}
