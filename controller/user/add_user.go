/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:33:51
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:23:32
 */
package user

import (
	"blogserver/library/response"
	svruser "blogserver/model/page/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type CreateUserParams struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

func AddUser(ctx *gin.Context) {
	params, errParams := checkAddUserParams(ctx)
	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	User := svruser.User{
		Username: params.Username,
		Password: params.Password,
		Email:    params.Email,
	}

	_, errAdd := User.AddUser(ctx)
	if errAdd != nil {
		log.Println("err", errAdd)
		response.Error(ctx, 200, errAdd.Error())
		return
	}

	response.Json(ctx, nil)
}

func checkAddUserParams(ctx *gin.Context) (*CreateUserParams, error) {
	res := &CreateUserParams{}
	err := ctx.ShouldBindJSON(res)

	if err != nil {
		log.Println("err", err)
		return nil, fmt.Errorf("params is invalid")
	}

	if len(res.Password) < 8 {
		msg := "password len is less than 8"
		log.Println("err", msg)
		return nil, fmt.Errorf(msg)
	}

	return res, nil
}
