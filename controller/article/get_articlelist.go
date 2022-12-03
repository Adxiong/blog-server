/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:37:38
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-03 23:23:00
 */
package article

import (
	"blogserver/library/response"
	svrarticle "blogserver/model/page/article"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type GetArticleListParams struct {
	Pn       int    `form:"pn"`
	Num      int    `form:"num"`
	Title    string `form:"title"`
	Content  string `form:"content"`
	AuthorID uint64 `form:"author_id"`
}

func GetArticleList(ctx *gin.Context) {
	params, errParams := checkGetArticleListParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Code(ctx, 205, errParams.Error())
		return
	}

	Article := svrarticle.Article{}
	cond := svrarticle.QueryArticleParam{}

	if len(params.Title) > 1 {
		cond.Title = &params.Title
	}

	if len(params.Content) > 1 {
		cond.Context = &params.Content
	}

	if params.AuthorID > 0 {
		cond.AuthorID = &params.AuthorID
	}

	articleList, errArticleList := Article.FindArticleList(ctx, params.Pn, params.Num, cond)
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
