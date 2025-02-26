package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserData(c *gin.Context, s database.Service) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get User id",
		})
		return
	}

	user, err := s.GetUserById(id.(uint))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to get user data" + err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"firstname":  user.FirstName,
		"surname":    user.Surname,
		"username":   user.Username,
		"email":      user.Email,
		"updated_at": user.UpdatedAt,
		"created_at": user.CreatedAt,
	})

}

func UpdateUserData(c *gin.Context, s database.Service) {
	id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to get user id",
		})
		return
	}

	var userData entity.User
	userData.ID = id.(uint)

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to bind JSON" + err.Error(),
		})
		return
	}

	if err := s.UpdateUser(userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to update user data" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated Successfully updated",
	})
}
