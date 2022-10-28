/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-29 00:21:07
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-29 00:40:15
 */
package article

import (
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
		ctx.JSON(200, gin.H{"msg": errParams.Error()})
		return
	}

	Article := svrarticle.Article{}
	articleDetail, errArticleDetail := Article.FindArticleByAID(ctx, params.Aid)
	if errArticleDetail != nil {
		log.Println("err", errArticleDetail)
		ctx.JSON(200, gin.H{"msg": errArticleDetail.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success", "data": articleDetail})

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
