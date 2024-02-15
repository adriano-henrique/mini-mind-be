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
			"message": "Nuggets fetched successfully",
			"status":  "success",
			"data":    nuggets,
		},
	)
}

func GetNuggetById(c *gin.Context) {
	db := database.DatabaseConnect()

	nuggetID := c.Param("id")
	if nuggetID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Startup ID is required"})
		return
	}

	nugget, err := models.FetchNugget(nuggetID, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Nugget fetched successfully",
			"status":  "success",
			"data":    nugget,
		},
	)
}

func CreateNugget(c *gin.Context) {
	db := database.DatabaseConnect()
	var input models.Nugget
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			},
		)
	}

	savedNugget, err := input.Save(db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			},
		)
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Nugget created successfully",
			"status":  "success",
			"data":    savedNugget,
		},
	)
}
