/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:59:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-05-02 16:13:23
 */
package link

import (
	"blogserver/library/response"
	"blogserver/model/page/link"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type getLinkListParams struct {
	ParentCID uint64 `json:"parent_cid"`
}

type respLink struct {
	ID         uint64     `json:"id"`          // 自增id
	Name       string     `json:"name"`        // link名称
	Url        string     `json:"url"`         // link地址
	LinkID     uint64     `json:"link_id"`     // link_id
	CategoryID uint64     `json:"category_id"` // link所在类别id
	Pic        string     `json:"pic"`         // 图片地址
	Brief      string     `json:"brief"`       // 简介
	CreateAt   *time.Time `json:"createAt"`    // 创建时间
	UpdateAt   *time.Time `json:"updateAt"`    // 更新时间
}
type respLinkList struct {
	Cid      uint64      `json:"cid"`
	Cname    string      `json:"category_name"`
	CreateAt *time.Time  `json:"create_at"`
	UpdateAt *time.Time  `json:"update_at"`
	List     []*respLink `json:"list"`
}

type respGetLinkList struct {
	List []*respLinkList `json:"list"`
}

func GetLinkList(ctx *gin.Context) {
	params, errParams := parseGetLinkList(ctx)

	if errParams != nil {
		fmt.Println("err", errParams)
		response.Error(ctx, 200, "参数解析错误")
	}

	svrCate := &link.LinkCategory{}
	svrLink := &link.Link{}

	// 指定了分类
	cateList := make([]*link.LinkCategory, 0)

	if params.ParentCID > 0 {
		cateInfo, errCateInfo := svrCate.GetLinkCategoryByCid(ctx, params.ParentCID)

		if errCateInfo != nil {
			fmt.Println("err", errCateInfo)
			response.Error(ctx, 200, "查询类别失败")
		}

		cateList = append(cateList, cateInfo)
	} else {
		lcList, errLcList := svrCate.GetLinkCategoryList(ctx)
		if errLcList != nil {
			fmt.Println("err", errLcList)
			response.Error(ctx, 200, "查询失败")
		}
		cateList = append(cateList, lcList...)
	}

	resList := make([]*respLinkList, len(cateList))

	for index, cate := range cateList {
		temp := &respLinkList{
			Cid:      cate.Cid,
			Cname:    cate.Name,
			CreateAt: cate.CreateAt,
			UpdateAt: cate.UpdateAt,
			List:     make([]*respLink, 0),
		}

		links, errLinks := svrLink.GetLinkListByCid(ctx, cate.Cid)

		if errLinks != nil {
			fmt.Println("err", errLinks)
		}

		tempList := make([]*respLink, len(links))
		for lnIndex, lk := range links {
			tempLink := &respLink{
				ID:         lk.ID,
				Name:       lk.Name,
				Url:        lk.Url,
				LinkID:     lk.LinkID,
				CategoryID: lk.CategoryId,
				Pic:        lk.Pic,
				Brief:      lk.Brief,
				CreateAt:   lk.CreateAt,
				UpdateAt:   lk.UpdateAt,
			}

			tempList[lnIndex] = tempLink
		}

		temp.List = tempList
		resList[index] = temp
	}

	result := &respGetLinkList{
		List: resList,
	}

	response.Json(ctx, result)
}

func parseGetLinkList(ctx *gin.Context) (*getLinkListParams, error) {
	p := &getLinkListParams{}

	err := ctx.ShouldBind(p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
