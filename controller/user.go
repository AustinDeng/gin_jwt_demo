package controller

import (
	"gin_jwt_demo/model"
	"gin_jwt_demo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录逻辑
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	data := make(map[string]interface{})
	// 查数据库
	msg, ok := model.CheckUser(username, password)

	if ok {
		token, err := utils.GenerateToken(username)
		if err != nil {
			msg = "生成 token 时出错"
		} else {
			data["token"] = token
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}

func GetUser(c *gin.Context) {
	username := c.Query("username")
	user := model.FindUserByUsername(username)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "OK",
		"user": user,
	})
}

func NewUser(c *gin.Context) {
	// 缺少校验用户名是否重复
	username := c.PostForm("username")
	password := c.PostForm("password")

	msg := model.CreateNewUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func ChangePassword(c *gin.Context) {
	// 缺少校验用户名是否存在
	username := c.Param("username")
	password := c.PostForm("password")

	msg := model.ChangeUserPassword(username, password)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func DeleteUser(c *gin.Context) {
	username := c.Query("username")

	msg := model.DeleteUserByUsername(username)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
