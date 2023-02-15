package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/controller"
	"github.com/gin-qt-business/app/middleware"
)

func Router(r *gin.Engine) {
	group := r.Group("/business")

	// 区分模块
	stockGroup := group.Group("/stock_info") // 行情模块
	newsGroup := group.Group("/news_info")   // 资讯新闻模块
	userGroup := group.Group("/user")        // 用户模块

	stock := controller.Stock{}
	news := controller.News{}
	user := controller.User{}

	// 全局加载日志中间件
	group.Use(middleware.RecordPostLog())

	userGroup.POST("/loginByPassword", func(ctx *gin.Context) {
		ctx.JSON(200, user.LoginByPassword(ctx))
	})

	stockGroup.POST("/getStockList", func(ctx *gin.Context) {
		ctx.JSON(200, stock.GetList(ctx))
	})

	newsGroup.POST("/getNewsList", func(ctx *gin.Context) {
		ctx.JSON(200, news.GetList(ctx))
	})
}
