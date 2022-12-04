/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:36:45
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:25:26
 */
package user

import (
	"fmt"
	"log"

	"blogserver/library/response"
	svruser "blogserver/model/page/user"

	"github.com/gin-gonic/gin"
)

type GetUserListParams struct {
	Pn  int `form:"pn"`
	Num int `form:"num"`
}

func GetUserList(ctx *gin.Context) {
	params, errParams := checkGetUserListParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	User := svruser.User{}
	userlist, errUserlist := User.FindUserList(ctx, params.Pn, params.Num, nil)
	if errUserlist != nil {
		log.Println("err", errUserlist)
		response.Error(ctx, 200, errUserlist.Error())
		return
	}

	response.Json(ctx, userlist)

}

func checkGetUserListParams(ctx *gin.Context) (*GetUserListParams, error) {
	res := &GetUserListParams{}

	err := ctx.ShouldBindJSON(res)

	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	if res.Pn < 0 {
		res.Pn = 0
	}

	if res.Num <= 0 {
		res.Num = 20
	}

	return res, nil
}
