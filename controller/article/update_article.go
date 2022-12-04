/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:37:54
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-12-04 16:22:28
 */
package article

import (
	"blogserver/library/response"
	svrarticle "blogserver/model/page/article"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type UpdateArticleParams struct {
	Title   string `form:"title"`
	Content string `form:"content"`
}

func UpdateArticle(ctx *gin.Context) {
	params, errParams := checkUpdateArticleParams(ctx)

	if errParams != nil {
		log.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
		return
	}

	Article := svrarticle.Article{}
	updateParams := &svrarticle.UpdateArticleParams{
		Title:   &params.Title,
		Content: &params.Content,
	}

	aid := uint64(22)
	errUpdate := Article.UpdateArticleByAID(ctx, aid, updateParams)

	if errUpdate != nil {
		log.Println("err", errUpdate)
		response.Error(ctx, 200, "user update failed")
		return
	}
	response.Json(ctx, nil)

}

func checkUpdateArticleParams(ctx *gin.Context) (*UpdateArticleParams, error) {
	res := &UpdateArticleParams{}

	err := ctx.ShouldBindJSON(res)
	if err != nil {
		msg := fmt.Errorf("params is invalid")
		log.Println("err", msg)
		return nil, msg
	}

	return res, nil
}
