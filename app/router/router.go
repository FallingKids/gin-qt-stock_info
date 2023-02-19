package router

import (
	"net/http"

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

	stockController := controller.StockController{}
	newsController := controller.NewsController{}
	userController := controller.UserController{}

	// 全局加载日志中间件
	group.Use(middleware.RecordPostLog())

	userGroup.POST("/register", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.Register(ctx))
	})

	userGroup.POST("/loginByPassword", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, userController.LoginByPassword(ctx))
	})

	// 校验登录信息
	group.Use(middleware.AccessCheck())

	stockGroup.POST("/getStockList", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, stockController.GetList(ctx))
	})

	newsGroup.POST("/getNewsList", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, newsController.GetList(ctx))
	})
}
