/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:56:53
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-15 23:51:22
 */
package db

import (
	"context"
	"fmt"
	"log"
)

func NewLinkCategory() *LinkCategory {
	return &LinkCategory{}
}

func NewLinkCategoryList() *LinkCategoryList {
	return &LinkCategoryList{}
}

// GetLinkCategoryList 获取所有link类别
func (lc *LinkCategory) GetLinkCategoryList() (*LinkCategoryList, error) {
	linkCategoryList := NewLinkCategoryList()
	res := GlobalDb.Table(lc.TableName()).Order("create_at aes").Find(linkCategoryList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return linkCategoryList, nil
}

// AddLinkCategory 创建一个link类别
func (lc *LinkCategory) AddLinkCategory(ctx context.Context) (*LinkCategory, error) {
	linkCategory := NewLinkCategory()

	res := GlobalDb.Table(lc.TableName()).Create(lc)

	if res.Error != nil {
		fmt.Println("err", res.Error)
		return nil, res.Error
	}

	return linkCategory, nil
}

// UpdateLinkCategory 更新link类别名称
func (lc *LinkCategory) UpdateLinkCategory(ctx context.Context, lc_id uint64, name string) (int64, error) {
	res := GlobalDb.Table(lc.TableName()).Where("cid = ?", lc_id).Update(LinkCategoryColumns.Name, name)

	if res.Error != nil {
		fmt.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// DelLinkCategory 根据cid删除分类
func (lc *LinkCategory) DelLinkCategory(ctx context.Context, lc_id uint64) (int64, error) {
	res := GlobalDb.Table(lc.TableName()).Where("cid = ?", lc_id).Update(LinkCategoryColumns.IsDel, IS_DEL)

	if res.Error != nil {
		fmt.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}
