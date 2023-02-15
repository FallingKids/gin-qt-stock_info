package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
	"github.com/gin-qt-business/app/logic"
	"github.com/gin-qt-business/app/type/data"
)

type User struct{}

func (u *User) LoginByPassword(ctx *gin.Context) *base.Response {
	return base.NewResponseSuccess(1)
}

func (u *User) Register(ctx *gin.Context) *base.Response {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return base.NewResponseError(err)
	}

	var input data.RegisterReq
	if err = json.Unmarshal(body, &input); err != nil {
		return base.NewResponseError(err)
	}

	logicUser := logic.User{}
	res, err := logicUser.Register(input)
	if err != nil {
		return base.NewResponseError(err)
	}

	return base.NewResponseSuccess(res)
}
