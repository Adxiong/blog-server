/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:58:11
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:13:01
 */
package link

import (
	"context"
	"fmt"
	"log"

	"blogserver/model/dao/db"
)

// AddLinkParams
type AddLinkParams struct {
	Name       string `json:"name"`        // link名称
	Url        string `json:"link"`        // 地址
	Pic        string `json:"pic"`         // 图片地址
	Brief      string `json:"brief"`       // 简介
	CategoryID uint64 `json:"category_id"` // 所在类别id
}

func (lk *Link) AddLink(ctx context.Context, params *AddLinkParams) (*Link, error) {

	dblink := db.Link{
		Name:       params.Name,
		Pic:        params.Pic,
		Brief:      params.Brief,
		Url:        params.Url,
		CategoryID: params.CategoryID,
	}

	dbResult, errDb := dblink.AddLink(ctx)

	if errDb != nil {
		err := fmt.Errorf("SERVICE_LINK_AddLink_Failed")
		log.Println("err", errDb)
		return nil, err
	}

	res := &Link{
		ID:         dbResult.ID,
		LinkID:     dbResult.LinkID,
		Name:       dbResult.Name,
		Pic:        dbResult.Pic,
		CategoryId: dbResult.CategoryID,
		CreateAt:   dbResult.CreateAt,
		UpdateAt:   dbResult.UpdateAt,
	}

	return res, nil
}
