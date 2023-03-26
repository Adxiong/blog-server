/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:56:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:02:25
 */
package link

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
)

func (lc *LinkCategory) GetLinkCategoryList(ctx context.Context) ([]*LinkCategory, error) {
	result := make([]*LinkCategory, 0)

	dbLink := db.NewLinkCategory()
	dbResult, dbErr := dbLink.GetLinkCategoryList(ctx)
	if dbErr != nil {
		err := fmt.Errorf("SERVICE_LINK_GetLinkCategoryList_Call_Db_GetLinkCategoryList_Failed")
		fmt.Println("err", dbErr)
		return nil, err
	}

	for _, item := range *dbResult {
		tmp := &LinkCategory{
			ID:       item.ID,
			Cid:      item.Cid,
			Name:     item.Name,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
		}

		result = append(result, tmp)
	}

	return result, nil
}
