package controller

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
	"github.com/gin-qt-business/app/constant"
	"github.com/gin-qt-business/app/logic"
	"github.com/gin-qt-business/app/type/data"
	"github.com/gin-qt-business/app/utils"
)

type UserController struct{}

func (u *UserController) LoginByPassword(ctx *gin.Context) *base.Response {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return base.NewResponseError(constant.READALL_ERROR)
	}

	var input data.LoginByPasswordReq
	if err = json.Unmarshal(body, &input); err != nil {
		return base.NewResponseError(constant.JSON_UNMARSHAL_ERROR)
	}

	logicUser := logic.UserLogic{
		LoginByPasswordParam: input,
		Ctx:                  ctx,
	}
	res, err := logicUser.LoginByPassword()
	if err != nil {
		return base.NewResponseError(constant.Message(err.Error()))
	}

	return base.NewResponseSuccess(res)
}

func (u *UserController) Register(ctx *gin.Context) *base.Response {
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		return base.NewResponseError(constant.READALL_ERROR)
	}

	var input data.RegisterReq
	if err = json.Unmarshal(body, &input); err != nil {
		return base.NewResponseError(constant.JSON_UNMARSHAL_ERROR)
	}

	if !utils.CheckMobile(input.Phone) {
		return base.NewResponseError(constant.PARAMS_VALIDATE_ERROR)
	}

	logicUser := logic.UserLogic{
		RegisterReqParams: input,
		Ctx:               ctx,
	}
	res, err := logicUser.Register()
	if err != nil {
		return base.NewResponseError(constant.Message(err.Error()))
	}

	return base.NewResponseSuccess(res)
}
