/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:36:19
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:24:28
 */
package user

import (
	"blogserver/library/response"
	svruser "blogserver/model/page/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type DelUserParams struct {
	UID uint64 `form:"uid" binding:"required"`
}

func DelUser(ctx *gin.Context) {
	params, errParams := checkDelUserParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	User := svruser.User{}
	errDel := User.DeleteUserByUID(ctx, params.UID)

	if errDel != nil {
		log.Println("err", errDel)
		response.Error(ctx, 200, "del account failed")
		return
	}

	response.Json(ctx, nil)
}

func checkDelUserParams(ctx *gin.Context) (*DelUserParams, error) {
	res := &DelUserParams{}

	err := ctx.ShouldBindJSON(res)

	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	if res.UID < 1 {
		msg := fmt.Errorf("uid is invalid")
		log.Println("err", msg)
		return nil, msg
	}
	return res, nil
}
