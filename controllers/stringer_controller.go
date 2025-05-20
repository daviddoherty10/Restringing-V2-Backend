package controllers

import (
	"Restringing-V2/entity"
	"Restringing-V2/internal/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StringSuggestion struct {
	Age      string `json:"age" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	Level    string `json:"level" binding:"required"`
	Injuries bool   `json:"injuries" binding:"required"`
}

func CreateApplication(c *gin.Context, db database.Service) {
	var potentialStringinger entity.PotentialStringer
	err := c.ShouldBindJSON(&potentialStringinger)
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get User id",
		})
	}

	potentialStringinger.ID = id.(uint)

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
		log.Default().Println("Failed to save Potential stringer to Database: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "There was an error adding user to application table" + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "accepted",
	})

}

func GetSuggestedStringSetup(c *gin.Context, s database.Service) {
	var stringSuggestion StringSuggestion

	var Strings []entity.String
	var PotenialStrings []entity.String

	if err := c.ShouldBindJSON(&stringSuggestion); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to Bind Json" + err.Error(),
		})
		return
	}

	if stringSuggestion.Sex == "male" {
		if stringSuggestion.Age == "<12" {
			for i := 0; i < len(Strings); i++ {
				if Strings[i].Type == "Synthetic Gut" {
					PotenialStrings = append(PotenialStrings, Strings[i])
				}
			}
		} else if stringSuggestion.Age == "<16" {
			for i := 0; i < len(Strings); i++ {
				if Strings[i].Type == "Synthetic Gut" {
					PotenialStrings = append(PotenialStrings, Strings[i])
				} else if Strings[i].Type == "Polyester" {
					PotenialStrings = append(PotenialStrings, Strings[i])
				}
			}
		} else if stringSuggestion.Age == ">16" {
			for i := 0; i < len(Strings); i++ {
				if Strings[i].Type == "Synthetic Gut" {
					PotenialStrings = append(PotenialStrings, Strings[i])
				} else if Strings[i].Type == "Polyester" {
					PotenialStrings = append(PotenialStrings, Strings[i])
				}
			}

		} else if stringSuggestion.Age == "<60" {
		}
	} else if stringSuggestion.Sex == "female" {
		if stringSuggestion.Age == "<12" {

		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "Invalid response",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"Potential Strings": "",
	})

}
