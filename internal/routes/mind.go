package routes

import (
	"github.com/adriano-henrique/mini-mind-be/internal/handler"
	"github.com/gin-gonic/gin"
)

func mindGroupRouter(baseRouter *gin.RouterGroup) {
	mind := baseRouter.Group("/mind")

	mind.GET("/get/:user_id", handler.GetMind)
}
