/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-19 22:17:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:12:24
 */
package link

import "time"

// Link
type Link struct {
	ID         uint64     `json:"id"`          // 自增id
	Name       string     `json:"name"`        // link名称
	Url        string     `json:"url"`         // link地址
	LinkID     uint64     `json:"link_id"`     // link_id
	CategoryId uint64     `json:"category_id"` // link所在类别id
	Pic        string     `json:"pic"`         // 图片地址
	Brief      string     `json:"brief"`       // 简介
	CreateAt   *time.Time `json:"createAt"`    // 创建时间
	UpdateAt   *time.Time `json:"updateAt"`    // 更新时间

}

type LinkCategory struct {
	ID       uint64     `json:"id"`       // 自增id
	Cid      uint64     `json:"cid"`      // 业务名称id
	Name     string     `json:"name"`     // 分类名称
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}
