package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateApplication(c *gin.Context, db database.Service) {
	var potentialStringinger entity.PotentialStringer
	err := c.ShouldBindJSON(&potentialStringinger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err = db.GetUserById(potentialStringinger.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There was an error loading user data" + err.Error(),
		})
		return
	}

	err = db.CreatePotenialStringer(potentialStringinger)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There was an error adding user to application table" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "accepted",
	})

}
