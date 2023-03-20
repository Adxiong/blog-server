/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-15 23:58:45
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-20 23:29:49
 */
package link

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
)

func (lk *LinkCategory) DelLinkCategory(ctx context.Context, linkCategoryID uint64) (bool, error) {
	dblc := db.NewLinkCategory()
	dbResult, dbErr := dblc.DelLinkCategory(ctx, linkCategoryID)

	if dbErr != nil {
		err := fmt.Errorf("SERVICE_LINK_DelLinkCategory_Call_Db_DelLinkCategory_Failed")
		fmt.Println("err", dbErr)
		return false, err
	}

	fmt.Printf("SERVICE_LINK_DelLinkCategory_Call_Db_DelLinkCategory_Affected:%d\n", dbResult)

	return true, nil
}
