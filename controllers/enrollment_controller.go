package controllers

import (
	"gymberapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEnrollments(c *gin.Context) {
	var enrollments []models.Enrollment
	if err := db.Preload("Member").Preload("Class").Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, enrollments)
}

func CreateEnrollment(c *gin.Context) {
	var enrollment models.Enrollment
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, enrollment)
}

func UpdateEnrollment(c *gin.Context) {
	var enrollment models.Enrollment
	if err := db.First(&enrollment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&enrollment)
	c.JSON(http.StatusOK, enrollment)
}

func DeleteEnrollment(c *gin.Context) {
	if err := db.Delete(&models.Enrollment{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
