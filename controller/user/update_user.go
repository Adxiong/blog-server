/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:36:32
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 23:40:34
 */
package user

import (
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
		ctx.JSON(200, gin.H{"msg": errParams})
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
		ctx.JSON(200, gin.H{"msg": "user update failed"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success"})

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
