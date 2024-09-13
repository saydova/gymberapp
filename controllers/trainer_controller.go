package controllers

import (
	"gymberapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTrainers(c *gin.Context) {
	var trainers []models.Trainer
	if err := db.Find(&trainers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, trainers)
}

func CreateTrainer(c *gin.Context) {
	var trainer models.Trainer
	if err := c.ShouldBindJSON(&trainer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&trainer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, trainer)
}

func UpdateTrainer(c *gin.Context) {
	var trainer models.Trainer
	if err := db.First(&trainer, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}
	if err := c.ShouldBindJSON(&trainer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&trainer)
	c.JSON(http.StatusOK, trainer)
}

func DeleteTrainer(c *gin.Context) {
	if err := db.Delete(&models.Trainer{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Trainer not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
