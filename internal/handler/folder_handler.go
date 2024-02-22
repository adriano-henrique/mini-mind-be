package handler

import (
	"fmt"
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

func DeleteFolder(c *gin.Context) {
	db := database.DatabaseConnect()

	folderID := c.Param("id")
	if folderID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": "FolderID is required",
				"data":    nil,
			},
		)
		return
	}

	err := models.DeleteFolder(folderID, db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
				"status":  "failed",
				"data":    nil,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Folder deleted successfully",
			"data":    nil,
		},
	)
}

func GetFolderByID(c *gin.Context) {
	db := database.DatabaseConnect()

	folderID := c.Param("id")
	if folderID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": "Folder ID is required",
				"data":    nil,
			},
		)
		return
	}

	folder, err := models.FetchFolder(folderID, db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
				"status":  "failed",
				"data":    folder,
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Folder fetched successfully",
			"status":  "success",
			"data":    folder,
		},
	)
}

func UpdateFolder(c *gin.Context) {
	db := database.DatabaseConnect()

	var updatedFolder *models.Folder
	if err := c.ShouldBindJSON(&updatedFolder); err != nil {
		fmt.Println("Failed to bind JSON")
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
				"status":  "failed",
				"data":    nil,
			},
		)
		return
	}

	updatedFolder, err := updatedFolder.UpdateFolder(db)
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
			"status":  "success",
			"message": "Nugget updated successfully",
			"data":    updatedFolder,
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
