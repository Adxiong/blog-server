/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:37:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:20:25
 */
package article

import (
	"blogserver/library/response"
	svrarticle "blogserver/model/page/article"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type DelArticleParams struct {
	AID uint64 `form:"aid" binding:"required"`
}

func DelArticle(ctx *gin.Context) {
	params, errParams := checkDelArticleParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	Article := svrarticle.Article{}
	errDel := Article.DeleteArticleByAID(ctx, params.AID)

	if errDel != nil {
		log.Println("err", errDel)
		response.Error(ctx, 200, "del account failed")
		return
	}
	response.Json(ctx, nil)
}

func checkDelArticleParams(ctx *gin.Context) (*DelArticleParams, error) {
	res := &DelArticleParams{}

	err := ctx.ShouldBindJSON(res)

	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	if res.AID < 1 {
		msg := fmt.Errorf("uid is invalid")
		log.Println("err", msg)
		return nil, msg
	}
	return res, nil
}
