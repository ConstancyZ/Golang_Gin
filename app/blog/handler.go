package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)
// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func login(c *gin.Context) {
	var login Login
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&login); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Pssword != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func likeBlog(c *gin.Context) {
	fmt.Println("like Blog")
}
