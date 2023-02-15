package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
)

type News struct{}

func (n *News) GetList(ctx *gin.Context) *base.Response {
	return nil
}
