/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-26 22:56:28
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:00:15
 */
package link

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
)

func (lc *LinkCategory) GetLinkCategoryByCid(ctx context.Context, cid uint64) (*LinkCategory, error) {
	result := &LinkCategory{}

	dbLink := db.NewLinkCategory()
	dbResult, dbErr := dbLink.GetLinkCategoryByCid(ctx, cid)
	if dbErr != nil {
		err := fmt.Errorf("SERVICE_LINK_GetLinkCategoryList_Call_Db_GetLinkCategoryList_Failed")
		fmt.Println("err", dbErr)
		return nil, err
	}

	result.ID = dbResult.ID
	result.Cid = dbLink.Cid
	result.Name = dbLink.Name
	result.CreateAt = dbLink.CreateAt
	result.UpdateAt = dbLink.UpdateAt

	return result, nil
}
