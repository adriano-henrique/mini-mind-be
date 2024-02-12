package handler

import (
	"net/http"

	"github.com/adriano-henrique/mini-mind-be/internal/database"
	"github.com/adriano-henrique/mini-mind-be/internal/models"
	"github.com/gin-gonic/gin"
)

func GetAllNuggets(c *gin.Context) {
	db := database.DatabaseConnect()

	nuggets, err := models.FetchAllNuggets(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Nuggets fetched succesfully",
			"status":  "success",
			"data":    nuggets,
		},
	)
}
