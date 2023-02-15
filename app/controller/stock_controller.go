package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
)

type Stock struct{}

func (s *Stock) GetList(ctx *gin.Context) *base.Response {
	return nil
}
