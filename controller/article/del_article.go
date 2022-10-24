/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:37:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 22:58:38
 */
package article

import (
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
		ctx.JSON(200, gin.H{"msg": errParams.Error()})
		return
	}

	Article := svrarticle.Article{}
	errDel := Article.DeleteArticleByAID(ctx, params.AID)

	if errDel != nil {
		log.Println("err", errDel)
		ctx.JSON(200, gin.H{"msg": "del account failed"})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success"})
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
