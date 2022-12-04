/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:33:57
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:19:37
 */
package article

import (
	"fmt"
	"log"

	"blogserver/library/response"
	svrarticle "blogserver/model/page/article"

	"github.com/gin-gonic/gin"
)

type CreateArticleParams struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func AddArticle(ctx *gin.Context) {
	params, errParams := checkAddArticleParams(ctx)
	if errParams != nil {
		log.Println("err", errParams)
		log.Printf("params:%+v", params)
		response.Error(ctx, 200, errParams.Error())
		return
	}
	aid := uint64(1231)
	Article := svrarticle.Article{
		Title:    params.Title,
		Content:  params.Content,
		AuthorID: aid,
	}

	_, errAdd := Article.AddArticle(ctx)
	if errAdd != nil {
		log.Println("err", errAdd)
		response.Error(ctx, 200, errParams.Error())
		return
	}
	response.Json(ctx, nil)
}

func checkAddArticleParams(ctx *gin.Context) (*CreateArticleParams, error) {
	res := &CreateArticleParams{}
	err := ctx.ShouldBind(res)

	if err != nil {
		log.Println("err", err)
		return nil, fmt.Errorf("params is invalid")
	}

	if len(res.Title) > 255 {
		msg := "title len is greater than 255"
		log.Println("err", msg)
		return nil, fmt.Errorf(msg)
	}
	return res, nil
}
