/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:59:54
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:14:41
 */
package link

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"blogserver/library/response"
	"blogserver/model/page/link"
)

type AddLinkParams struct {
	Name       string `json:"name" binding:"required"`
	Url        string `json:"url" binding:"required"`
	Pic        string `json:"icon" binding:"required"`
	Brief      string `json:"brief" binding:"required"`
	CategoryID uint64 `json:"category_id" binding:"required"`
}

type AddLinkResponse struct {
	ID uint64 `json:"id"`
}

// AddLink 增加一条link
func AddLink(ctx *gin.Context) {
	params, errParams := parseAddLinkParams(ctx)

	if errParams != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, errParams.Error())
	}

	svrLink := &link.Link{}

	addLinkParams := &link.AddLinkParams{
		Name:       params.Name,
		Url:        params.Url,
		Pic:        params.Pic,
		Brief:      params.Brief,
		CategoryID: params.CategoryID,
	}

	res, errRes := svrLink.AddLink(ctx, addLinkParams)

	if errRes != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, errRes.Error())
	}

	result := &AddLinkResponse{
		ID: res.ID,
	}

	response.Json(ctx, result)
}

func parseAddLinkParams(ctx *gin.Context) (*AddLinkParams, error) {
	p := &AddLinkParams{}

	err := ctx.ShouldBindJSON(p)

	if err != nil {
		return p, err
	}

	return p, nil
}
