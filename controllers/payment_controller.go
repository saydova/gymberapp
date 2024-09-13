package controllers

import (
	"gymberapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPayments(c *gin.Context) {
	var payments []models.Payment
	if err := db.Preload("Member").Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func CreatePayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&payment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payment)
}

func UpdatePayment(c *gin.Context) {
	var payment models.Payment
	if err := db.First(&payment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&payment)
	c.JSON(http.StatusOK, payment)
}

func DeletePayment(c *gin.Context) {
	if err := db.Delete(&models.Payment{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
