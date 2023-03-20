/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:58:34
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-20 23:29:34
 */
package link

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
)

func (lk *Link) DelLink(ctx context.Context, linkID uint64) (bool, error) {
	dblink := db.NewLink()
	dbResult, dbErr := dblink.DelLink(ctx, linkID)

	if dbErr != nil {
		err := fmt.Errorf("SERVICE_LINK_DelLink_Call_Db_DelLink_Failed")
		fmt.Println("err", dbErr)
		return false, err
	}

	fmt.Printf("SERVICE_LINK_DelLink_Call_Db_DelLink_Affected:%d\n", dbResult)

	return true, nil
}
