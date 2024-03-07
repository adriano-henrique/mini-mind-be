package handler

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nugget ID is required"})
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

func GetNuggetByKey(c *gin.Context) {
	db := database.DatabaseConnect()

	nuggetKey := c.Param("key")
	if nuggetKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Nugget key is required"})
		return
	}

	nugget, err := models.FetchAllNuggetsByKey(nuggetKey, db)
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

func GetNuggetByFolderID(c *gin.Context) {
	db := database.DatabaseConnect()

	folderID := c.Param("folder_id")
	if folderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Folder ID is required"})
		return
	}

	nuggets, err := models.FetchNuggetsByFolderID(folderID, db)
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
		return
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
		return
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

func UpdateNugget(c *gin.Context) {
	db := database.DatabaseConnect()

	var updatedNugget *models.Nugget
	if err := c.ShouldBindJSON(&updatedNugget); err != nil {
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

	updatedNugget, err := updatedNugget.UpdateNugget(db)
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
			"data":    updatedNugget,
		},
	)
}

func DeleteNugget(c *gin.Context) {
	db := database.DatabaseConnect()
	nuggetID := c.Param("id")
	if nuggetID == "" {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "failed",
				"message": "NuggetID is required",
			},
		)
		return
	}

	err := models.DeleteNugget(nuggetID, db)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": err.Error(),
				"status":  "failed",
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":  "success",
			"message": "Nugget deleted successfully",
			"data":    nil,
		},
	)
}
