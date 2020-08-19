package controller

import (
	"github.com/gin-gonic/gin"
	"lish.org/ginBlogs/model"
	"lish.org/ginBlogs/repository"
	"lish.org/ginBlogs/response"
	"lish.org/ginBlogs/vo"
	"log"
	"strconv"
)

type ICategoryController interface {
	RestController
}

// 结构体可以引入接口生成接口函数： struct - implement - interface =》 func
type CategoryController struct {
	//DB *gorm.DB		// 引入连接池
	Repository repository.ICategoryRepository
}

func NewCategoryController() ICategoryController {
	categoryController := CategoryController{Repository: repository.NewCategoryRepository()}
	//db.AutoMigrate(model.Category{})
	//return CategoryController{DB:db}
	categoryController.Repository.(repository.CategoryRepository).DB.AutoMigrate(model.Category{})
	return categoryController
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	//ctx.BindJSON(&requestCategory)		// 获取json和text
	//ctx.ShouldBind(&requestCategory)		// 获取json
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		log.Println(err.Error())
		response.Fail(ctx, "数据验证错误", nil)
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		response.Fail(ctx, "数据验证错误", nil)
		return
	}
	response.Success(ctx, gin.H{"category": category}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	// 绑定 body 的参数
	var requestCategory vo.CreateCategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误", nil)
		log.Println(err.Error())
		return
	}
	// 获取 path 参数  ; strconv.Atoi类型转换 str --》 int
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var updateCategory model.Category
	//if c.DB.First(&updateCategory,categoryId).RecordNotFound(){
	//	response.Fail(ctx,"分类不存在",nil)
	//}
	//// 跟新分类 map struct
	//c.DB.Model(&updateCategory).Update("name",requestCategory.Name)
	var updateCategory *model.Category
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	updateCategory, err = c.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": updateCategory}, "更新成功")
}

func (c CategoryController) Show(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id")) // 获取url的路径参数
	//category, err := c.Repository
	//var category model.Category
	////ctx.ShouldBind(&category)	// 获取form-data参数
	//if c.DB.First(&category, categoryId).RecordNotFound() {
	//	response.Fail(ctx,"分类不存在",nil)
	//}
	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
	}
	response.Success(ctx, gin.H{"category": category}, "")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var category model.Category
	//if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
	//	response.Fail(ctx,"删除失败",nil)
	//	return
	//}
	c.Repository.DeleteById(categoryId)
	response.Success(ctx, nil, "删除成功")
}
