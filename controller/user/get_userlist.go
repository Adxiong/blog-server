/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:36:45
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 22:59:34
 */
package user

import (
	"fmt"
	"log"

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
		ctx.JSON(200, gin.H{"msg": errParams.Error()})
		return
	}

	User := svruser.User{}
	userlist, errUserlist := User.FindUserList(ctx, params.Pn, params.Num, nil)
	if errUserlist != nil {
		log.Println("err", errUserlist)
		ctx.JSON(200, gin.H{"msg": errUserlist.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success", "data": userlist})

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
