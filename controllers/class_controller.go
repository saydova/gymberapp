package controllers

import (
	"gymberapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClasses(c *gin.Context) {
	var classes []models.Class
	if err := db.Preload("Trainer").Find(&classes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, classes)
}

func CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, class)
}

func UpdateClass(c *gin.Context) {
	var class models.Class
	if err := db.First(&class, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	if err := c.ShouldBindJSON(&class); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&class)
	c.JSON(http.StatusOK, class)
}

func DeleteClass(c *gin.Context) {
	if err := db.Delete(&models.Class{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
