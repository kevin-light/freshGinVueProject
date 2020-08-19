package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"lish.org/ginBlogs/common"
	"lish.org/ginBlogs/controller"
	"lish.org/ginBlogs/model"
	"lish.org/ginBlogs/response"
	"strconv"
)

type ICategoryController interface {
	controller.RestController
}

// 结构体可以引入接口生成接口函数： struct - implement - interface =》 func
type CategoryController struct {
	DB *gorm.DB // 引入连接池
}

func NewCategoryController() ICategoryController {
	db := common.GetDB()
	db.AutoMigrate(model.Category{})
	return CategoryController{DB: db}

}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	//ctx.BindJSON(&requestCategory)		// 获取json和text
	ctx.ShouldBind(&requestCategory) // 获取json

	fmt.Println(requestCategory, "---0", requestCategory.Name)
	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误， 分类名称必填", nil)
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"category": requestCategory}, "")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定 body 的参数
	var requestCategory model.Category
	ctx.ShouldBind(&requestCategory)
	fmt.Println(requestCategory, "---2")
	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误， 分类名称必填", nil)
	}
	// 获取 path 参数  ; strconv.Atoi类型转换 str --》 int
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, "分类不存在", nil)
	}
	// 跟新分类 map struct
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id")) // 获取url的路径参数
	//category, err := c.Repository
	var category model.Category
	//ctx.ShouldBind(&category)	// 获取form-data参数

	fmt.Println(category, "---3")
	if c.DB.First(&category, categoryId).RecordNotFound() {
		response.Fail(ctx, "分类不存在", nil)
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var category model.Category
	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, "删除失败", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "")
}
