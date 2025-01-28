package user

import (
	"Restringing-V2/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleGetUserById(db database.Service, c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	user, err := db.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
