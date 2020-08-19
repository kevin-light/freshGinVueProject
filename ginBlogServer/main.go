package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lish.org/ginBlogs/common"
)

func main() {
	db := common.InitDb()
	defer db.Close()

	r := gin.Default()
	r = CollectRoute(r)

	panic(r.Run())
}
