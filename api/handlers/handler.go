package handlers

import (
	"net/http"
	"strconv"

	"github.com/Shivd131/api/db"
	"github.com/Shivd131/api/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// list of all items
func GetItems(c *gin.Context) {
	var items []models.Item
	db.GetDB().Find(&items)
	c.JSON(http.StatusOK, items)
}

// item by ID
func GetItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var item models.Item
	result := db.GetDB().First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving item"})
		}
		return
	}

	c.JSON(http.StatusOK, item)
}

// new item
func CreateItem(c *gin.Context) {
	var newItem models.Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.GetDB().Create(&newItem)
	c.JSON(http.StatusCreated, newItem)
}

// updates existing item
func UpdateItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	var updatedItem models.Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if item exists
	var existingItem models.Item
	result := db.GetDB().First(&existingItem, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving item"})
		}
		return
	}

	// Update
	db.GetDB().Model(&existingItem).Updates(updatedItem)
	c.JSON(http.StatusOK, existingItem)
}

func DeleteItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID"})
		return
	}

	// Check if item exists
	var item models.Item
	result := db.GetDB().First(&item, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving item"})
		}
		return
	}

	// Delete
	db.GetDB().Delete(&item)
	c.JSON(http.StatusNoContent, nil)
}
