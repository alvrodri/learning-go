package routes

import (
	"Twitter/main/models"
	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

func GetUser(context *gin.Context) {
	coll := mgm.Coll(&models.User{})
	result := []models.User{}
	err := coll.SimpleFind(&result, bson.M{
		"username": context.Param("username"),
	}, options.Find().SetProjection(bson.D{
		{"password", 0},
	}))

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, result)
}

type CreateUserForm struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func CreateUser(context *gin.Context) {
	var jsonForm CreateUserForm
	if err := context.ShouldBindJSON(&jsonForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &models.User{
		UserName: jsonForm.UserName,
		Password: jsonForm.Password,
	}

	coll := mgm.Coll(user)

	var result []models.User
	err := coll.SimpleFind(&result, bson.M{
		"username": user.UserName,
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(result) != 0 {
		context.JSON(http.StatusConflict, gin.H{
			"error": "Username is already taken",
		})
		return
	}

	err = coll.Create(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.String(http.StatusOK, "OK")
}
