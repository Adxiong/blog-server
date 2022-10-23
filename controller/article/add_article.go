/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:33:57
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 23:43:22
 */
package article

import (
	"fmt"
	"log"

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
		ctx.JSON(200, gin.H{"msg": errParams})
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
		ctx.JSON(200, gin.H{"msg": errAdd})
		return
	}
	ctx.JSON(200, gin.H{"msg": "success"})
}

func checkAddArticleParams(ctx *gin.Context) (*CreateArticleParams, error) {
	res := &CreateArticleParams{}
	err := ctx.ShouldBindJSON(res)

	if err != nil {
		log.Println("err", err)
		return nil, fmt.Errorf("params is invalid")
	}

	if len(res.Title) > 255 {
		msg := "password len is greater than 255"
		log.Println("err", msg)
		return nil, fmt.Errorf(msg)
	}

	return res, nil
}