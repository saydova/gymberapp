package controllers

import (
	"gymberapp/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetMembers(c *gin.Context) {
	var members []models.Member
	if err := db.Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, members)
}

func CreateMember(c *gin.Context) {
	var member models.Member
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := db.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, member)
}

func UpdateMember(c *gin.Context) {
	var member models.Member
	if err := db.First(&member, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&member)
	c.JSON(http.StatusOK, member)
}

func DeleteMember(c *gin.Context) {
	if err := db.Delete(&models.Member{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
