/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-16 00:00:57
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:08:45
 */
package link

import (
	"blogserver/library/response"
	"blogserver/model/page/link"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type getLinkCategoryListParams struct {
}

type respGetLinkCagegory struct {
	ID       uint64     `json:"id"`       // 自增id
	Cid      uint64     `json:"cid"`      // 业务名称id
	Name     string     `json:"name"`     // 分类名称
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}

func GetLinkCategoryList(ctx *gin.Context) {
	_, errParams := parseAddLinkCategoryParams(ctx)
	if errParams != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, "参数解析失败")
	}

	svrLC := link.LinkCategory{}
	res, err := svrLC.GetLinkCategoryList(ctx)

	if err != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, "查询失败")
	}

	result := make([]*respGetLinkCagegory, 0)

	for _, item := range res {
		temp := &respGetLinkCagegory{
			ID:       item.ID,
			Cid:      item.Cid,
			Name:     item.Name,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
		}

		result = append(result, temp)
	}

	response.Json(ctx, result)

}

func parseGetLinkCategoryList(ctx *gin.Context) (*getLinkCategoryListParams, error) {
	p := &getLinkCategoryListParams{}

	err := ctx.ShouldBindJSON(p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
