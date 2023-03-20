/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:56:50
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-20 23:39:39
 */
package link

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
)

func (link *Link) GetLinkListByCid(ctx context.Context) (*[]Link, error) {
	result := make([]Link, 0)

	dbLink := db.NewLink()
	dbResult, dbErr := dbLink.GetLinkList(ctx)
	if dbErr != nil {
		err := fmt.Errorf("SERVICE_LINK_GetLinkList_Call_Db_GetLinkList_Failed")
		fmt.Println("err", dbErr)
		return nil, err
	}

	for _, item := range *dbResult {
		tmp := Link{
			ID:       item.ID,
			LinkID:   item.LinkID,
			Pic:      item.Pic,
			Brief:    item.Brief,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
		}

		result = append(result, tmp)
	}

	return &result, nil
}
