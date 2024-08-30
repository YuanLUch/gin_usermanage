package api

import (
	"LoveDiary/middleware"
	"LoveDiary/model"
	"LoveDiary/services"
	"LoveDiary/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	var code int
	_ = c.ShouldBindJSON(&data)

	user, err := services.GetUserByUsername(data.UserName, data.Password)
	if err != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  err,
			"message": errmsg.GetErrMsg(err),
		})
	} else {
		token, code = middleware.SetToken(data.UserName, user.ID, user.Role)
		c.JSON(http.StatusOK, gin.H{
			"status":   code,
			"message":  errmsg.GetErrMsg(code),
			"username": user.UserName,
			"ID":       user.ID,
			"role":     user.Role,
			"token":    token,
		})
	}
}
