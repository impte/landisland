package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/impte/landisland/biz"
	"github.com/impte/landisland/reqres"
	"net/http"
)

func Login(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"message": "未知错误",
			})
		}
	}()
	var loginRequest reqres.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loginResponse := biz.Login(loginRequest)
	c.JSON(http.StatusOK, gin.H{"token": loginResponse.Token})
}

func Logout(c *gin.Context) {

}
