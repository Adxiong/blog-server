/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:37:38
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-24 23:21:47
 */
package article

import (
	svrarticle "blogserver/model/page/article"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type GetArticleListParams struct {
	Pn  int `form:"pn"`
	Num int `form:"num"`
}

func GetArticleList(ctx *gin.Context) {
	params, errParams := checkGetArticleListParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		ctx.JSON(200, gin.H{"msg": errParams.Error()})
		return
	}

	Article := svrarticle.Article{}
	articleList, errArticleList := Article.FindArticleList(ctx, params.Pn, params.Num, nil)
	if errArticleList != nil {
		log.Println("err", errArticleList)
		ctx.JSON(200, gin.H{"msg": errArticleList.Error()})
		return
	}

	ctx.JSON(200, gin.H{"msg": "success", "data": articleList})

}

func checkGetArticleListParams(ctx *gin.Context) (*GetArticleListParams, error) {
	res := &GetArticleListParams{}

	err := ctx.ShouldBind(res)

	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", err)
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
