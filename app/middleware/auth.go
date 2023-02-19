package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-qt-business/app/base"
	"github.com/gin-qt-business/app/constant"
	"github.com/gin-qt-business/app/utils"
)

const HEADER_TOKEN = "Authorization"

func AccessCheck() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get(HEADER_TOKEN)
		if token != "" {
			jwt := &utils.JWT{
				Token: token,
			}
			valid, err := jwt.ValidateToken()
			if err != nil {
				context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(constant.Message(err.Error())))
			}
			if !valid {
				context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(constant.TOKEN_INVALID_ERROR))
			}
		} else {
			context.AbortWithStatusJSON(http.StatusUnauthorized, base.NewResponseError(constant.TOKEN_INVALID_ERROR))
		}
		context.Next()
	}
}
