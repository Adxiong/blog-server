/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:57:01
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:11:22
 */
package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

func NewLink() *Link {
	return &Link{}
}

func NewLinkList() *LinkList {
	return &LinkList{}
}

// AddLink 增加link
func (link *Link) AddLink(ctx context.Context) (*Link, error) {
	err := GlobalDb.Table(link.TableName()).Create(link).Error
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return link, nil
}

// DelLink 删除link
func (link *Link) DelLink(ctx context.Context, link_id uint64) (int64, error) {
	res := GlobalDb.Table(link.TableName()).Where("link_id = ?", link_id).Update(LinkCategoryColumns.IsDel, IS_DEL)
	if res.Error != nil {
		log.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// 更新link
func (link *Link) UpdateLink(ctx context.Context, link_id uint64, values map[string]interface{}) (int64, error) {
	res := GlobalDb.Table(link.TableName()).Where("link_id  = ?", link_id).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// GetLinkList 获取所有的link列表
func (link *Link) GetLinkListByCid(ctx context.Context, cid uint64) (*LinkList, error) {
	linkList := NewLinkList()

	res := GlobalDb.Table(link.TableName()).Where("category_id =?", cid).Find(linkList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return linkList, nil
}

func (link *Link) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	link.CreateAt = &t
	link.UpdateAt = &t

	return nil
}

func (link *Link) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[LinkColumns.UpdateAt]; !ok {
			t := time.Now()
			values[LinkColumns.UpdateAt] = &t
		}

		if _, ok := values[LinkColumns.Version]; !ok {
			values[LinkColumns.UpdateAt] = gorm.Expr(fmt.Sprintf("%s + ?", LinkColumns.Version), 1)
		}
	}

	return nil
}
