/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:56:53
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-04-09 23:43:28
 */
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func NewLinkCategory() *LinkCategory {
	return &LinkCategory{}
}

func NewLinkCategoryList() *LinkCategoryList {
	return &LinkCategoryList{}
}

// GetLinkCategoryByCid 根据cid获取类别
func (lc *LinkCategory) GetLinkCategoryByCid(ctx context.Context, cid uint64) (*LinkCategory, error) {
	linkCategory := NewLinkCategory()
	res := GlobalDb.Table(lc.TableName()).Find(linkCategory)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}
	return linkCategory, nil
}

// GetLinkCategoryList 获取所有link类别
func (lc *LinkCategory) GetLinkCategoryList(ctx context.Context) (*LinkCategoryList, error) {
	linkCategoryList := NewLinkCategoryList()
	res := GlobalDb.Table(lc.TableName()).Order("create_at asc").Find(linkCategoryList)

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

func (lc *LinkCategory) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	lc.CreateAt = &t
	lc.UpdateAt = &t

	return nil
}

func (lc *LinkCategory) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[LinkCategoryColumns.UpdateAt]; !ok {
			t := time.Now()
			values[LinkCategoryColumns.UpdateAt] = &t
		}

		if _, ok := values[LinkCategoryColumns.Version]; !ok {
			values[LinkCategoryColumns.UpdateAt] = gorm.Expr(fmt.Sprintf("%s + ?", LinkCategoryColumns.Version), 1)
		}
	}

	return nil
}
