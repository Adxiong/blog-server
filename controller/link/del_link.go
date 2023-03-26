/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:59:04
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 16:45:00
 */
package link

import (
	"blogserver/library/response"
	"blogserver/model/page/link"
	"fmt"

	"github.com/gin-gonic/gin"
)

type delLinkParams struct {
	LinkID uint64 `json:"link_id" binding:"required"`
}

func DelLink(ctx *gin.Context) {
	param, errParam := parseDelLink(ctx)

	if errParam != nil {
		err := fmt.Errorf("参数解析失败")
		fmt.Println("err", errParam)
		response.Error(ctx, 200, err.Error())
	}

	svrLink := link.Link{}
	delRes, errDelRes := svrLink.DelLink(ctx, param.LinkID)

	if errDelRes != nil || !delRes {
		fmt.Println("err", errDelRes)
		response.Error(ctx, 200, "删除失败")
	}

	response.Json(ctx, struct{}{})

}

func parseDelLink(ctx *gin.Context) (*delLinkParams, error) {
	p := &delLinkParams{}

	err := ctx.BindJSON(p)
	if err != nil {
		return nil, err
	}

	return p, nil
}
