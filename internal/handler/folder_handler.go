package handler

import (
	"net/http"

	"github.com/adriano-henrique/mini-mind-be/internal/database"
	"github.com/adriano-henrique/mini-mind-be/internal/models"
	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {
	db := database.DatabaseConnect()
	var input models.Folder
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			},
		)
		return
	}

	savedFolder, err := input.Save(db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"message": "Folder created successfully",
			"status":  "success",
			"data":    savedFolder,
		},
	)
}

func GetAllFolders(c *gin.Context) {
	db := database.DatabaseConnect()
	folders, err := models.FetchAllFolders(db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": err.Error(),
				"data":    nil,
			},
		)
		return
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Folders fetched successfully",
			"status":  "success",
			"data":    folders,
		},
	)
}
