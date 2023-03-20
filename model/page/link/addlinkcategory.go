/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:57:48
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-20 23:19:09
 */
package link

import (
	"context"
	"fmt"

	"blogserver/model/dao/db"
)

func (lc *LinkCategory) AddLinkCategory(ctx context.Context, name string) (*LinkCategory, error) {
	dblc := &db.LinkCategory{
		Name: name,
	}
	dbres, dberr := dblc.AddLinkCategory(ctx)

	if dberr != nil {
		err := fmt.Errorf("SERVICE_LINK_AddCategoryLink_Failed")
		fmt.Println("err", dberr)
		return nil, err
	}

	res := &LinkCategory{
		ID:       dbres.ID,
		Cid:      dbres.Cid,
		Name:     dbres.Name,
		CreateAt: dblc.CreateAt,
		UpdateAt: dblc.UpdateAt,
	}

	return res, nil
}
