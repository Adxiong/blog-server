/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:36:32
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:26:27
 */
package user

import (
	"blogserver/library/response"
	svruser "blogserver/model/page/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type UpdateUserParams struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func UpdateUser(ctx *gin.Context) {
	params, errParams := checkUpdateUserParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	User := svruser.User{}
	updateParams := &svruser.UpdateUserParams{
		Username: &params.Username,
		Password: &params.Password,
	}

	uid := uint64(22)
	errUpdate := User.UpdateUserByUID(ctx, uid, updateParams)

	if errUpdate != nil {
		log.Println("err", errUpdate)
		response.Error(ctx, 200, "user update failed")
		return
	}

	response.Json(ctx, nil)

}

func checkUpdateUserParams(ctx *gin.Context) (*UpdateUserParams, error) {
	res := &UpdateUserParams{}

	err := ctx.ShouldBindJSON(res)
	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	return res, nil
}
