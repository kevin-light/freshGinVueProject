package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lish.org/ginBlogs/common"
	"lish.org/ginBlogs/model"
	"lish.org/ginBlogs/util"
	"log"
	"net/http"
)

func Register(c *gin.Context) {

	DB := common.GetDB()
	// 获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	fmt.Println("--1", len(telephone), telephone)
	// 数据验证
	if len(telephone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号必须为11位"})
		return
	}
	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码必须要大于六位"})
		return
	}
	// 如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
		fmt.Println("---name: ", name)
	}

	log.Println(name, telephone, password)
	// 判断手机号是否纯在
	if isTelephoneExist(DB, telephone) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户已存在"})
		return
	}
	// 创建用户
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	DB.Create(&newUser)
	// 返回结果
	c.JSON(200, gin.H{
		"msg:": "注册成功",
	})
}

func isTelephoneExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false

}
