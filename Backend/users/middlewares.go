package users

import (
	"net/http"

	"blog.com/common"
	"blog.com/models"
	"github.com/gin-gonic/gin"
)

func UpdateContextUserModel(c *gin.Context, my_user_id uint) {
	var myUserModel models.UserModel
	if my_user_id != 0 {
		db := common.GetDB()
		db.First(&myUserModel, my_user_id)
	}
	c.Set("my_user_id", my_user_id)
	c.Set("my_user_model", myUserModel)
}

func AuthMiddleware(auto401 bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		UpdateContextUserModel(c, 0)

		clientToken := common.BearerAuth(c)

		if clientToken == "" {
			return
		}

		claims, err := common.ValidateToken(clientToken)

		if err != "" {
			if auto401 {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				c.Abort()
			}
			return
		}
		UpdateContextUserModel(c, claims.Id)
	}
}
