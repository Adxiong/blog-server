/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-12-03 23:07:53
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-03 23:40:21
 */
package response

import (
	"github.com/gin-gonic/gin"
)

type responseBody struct {
	Errno  int    `json:"errno"`
	ErrMsg string `json:"errmsg"`
	Data   any    `json:"data"`
}

func Json(ctx *gin.Context, data any) {
	resultData := responseBody{
		Errno:  0,
		ErrMsg: "success",
		Data:   data,
	}
	ctx.JSON(200, resultData)
}

func Error(ctx *gin.Context, code int, msg string) {
	resultData := responseBody{
		Errno:  code,
		ErrMsg: msg,
		Data:   make([]any, 0),
	}
	ctx.JSON(200, resultData)
}
