/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:58:11
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-19 22:17:34
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
	Name  string `json:"name"`  // link名称
	Pic   string `json:"pic"`   // 图片地址
	Brief string `json:"brief"` // 简介
}

func (lk *Link) AddLink(ctx context.Context, params *AddLinkParams) (*Link, error) {

	dblink := db.Link{
		Name:  params.Name,
		Pic:   params.Pic,
		Brief: params.Brief,
	}

	dbResult, errDb := dblink.AddLink(ctx)

	if errDb != nil {
		err := fmt.Errorf("SERVICE_LINK_AddLink_Failed")
		log.Println("err", errDb)
		return nil, err
	}

	res := &Link{
		ID:       dbResult.ID,
		LinkID:   dbResult.LinkID,
		Name:     dbResult.Name,
		Pic:      dbResult.Pic,
		CreateAt: dbResult.CreateAt,
		UpdateAt: dbResult.UpdateAt,
	}

	return res, nil
}
