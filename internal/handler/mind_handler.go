package handler

import (
	"net/http"

	"github.com/adriano-henrique/mini-mind-be/internal/database"
	"github.com/adriano-henrique/mini-mind-be/internal/models"
	"github.com/gin-gonic/gin"
)

func GetMind(c *gin.Context) {
	db := database.DatabaseConnect()
	mind, err := models.FetchMind(c.Param("user_id"), db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Mind fetched successfully",
			"status":  "success",
			"data":    mind,
		},
	)
}
