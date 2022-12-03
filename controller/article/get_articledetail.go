/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-29 00:21:07
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-03 23:28:40
 */
package article

import (
	"blogserver/library/response"
	svrarticle "blogserver/model/page/article"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type GetArticleDetailParams struct {
	Aid uint64 `form:"aid" binding:"required"`
}

func GetArticleDetail(ctx *gin.Context) {
	params, errParams := checkGetArticleDetailParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Code(ctx, 205, errParams.Error())
		return
	}

	Article := svrarticle.Article{}
	articleDetail, errArticleDetail := Article.FindArticleByAID(ctx, params.Aid)
	if errArticleDetail != nil {
		log.Println("err", errArticleDetail)
		response.Code(ctx, 205, errParams.Error())
		return
	}

	response.Json(ctx, articleDetail)
}

func checkGetArticleDetailParams(ctx *gin.Context) (*GetArticleDetailParams, error) {
	res := &GetArticleDetailParams{}

	err := ctx.ShouldBindQuery(res)

	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", err)
		return nil, msg
	}

	if res.Aid < 0 {
		msg := fmt.Errorf("params aid is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	return res, nil
}
