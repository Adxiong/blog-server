/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:57:01
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 23:26:13
 */
package db

import (
	"context"
	"log"
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
func (link *Link) DelLink(ctx context.Context, link_id int64) (int64, error) {
	res := GlobalDb.Table(link.TableName()).Where("link_id = ?", link_id).Update(LinkCategoryColumns.IsDel, 1)
	if res.Error != nil {
		log.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// 更新link
func (link *Link) UpdateLink(ctx context.Context, link_id uint64, values map[string]interface{}) (int64, error) {
	res := GlobalDb.Table(link.TableName()).Where("link_id = ?", link_id).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return res.RowsAffected, res.Error
	}

	return res.RowsAffected, nil
}

// GetLinkList 获取所有的link列表
func (link *Link) GetLinkList(ctx context.Context) (*LinkList, error) {
	linkList := NewLinkList()

	res := GlobalDb.Table(link.TableName()).Order("create_at aes").Find(linkList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return linkList, nil
}
