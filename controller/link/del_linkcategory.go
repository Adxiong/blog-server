/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:59:16
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 16:51:43
 */
package link

import (
	"blogserver/library/response"
	"blogserver/model/page/link"
	"fmt"

	"github.com/gin-gonic/gin"
)

type delLinkCateGoryParams struct {
	LcID uint64 `json:"lc_id" binding:"required"`
}

func DelLinkCateGory(ctx *gin.Context) {
	params, errParams := parseDelLinkCategory(ctx)

	if errParams != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, "参数错误")
	}

	svrLC := link.LinkCategory{}
	delRes, errDelRes := svrLC.DelLinkCategory(ctx, params.LcID)
	if errDelRes != nil || !delRes {
		fmt.Println("err", errDelRes)
		response.Error(ctx, 200, "删除失败")
	}

	response.Json(ctx, struct{}{})

}

func parseDelLinkCategory(ctx *gin.Context) (*delLinkCateGoryParams, error) {
	p := &delLinkCateGoryParams{}
	err := ctx.ShouldBindJSON(ctx)

	if err != nil {
		return nil, err
	}

	return p, nil
}
