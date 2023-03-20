/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-19 22:17:26
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-20 22:40:19
 */
package link

import "time"

// Link
type Link struct {
	ID       uint       `json:"id"`       // 自增id
	Name     string     `json:"name"`     // link名称
	LinkID   uint       `json:"link_id"`  // link_id
	Pic      string     `json:"pic"`      // 图片地址
	Brief    string     `json:"brief"`    // 简介
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间

}

type LinkCategory struct {
	ID       uint       `json:"id"`       // 自增id
	Cid      uint       `json:"cid"`      // 业务名称id
	Name     string     `json:"name"`     // 分类名称
	CreateAt *time.Time `json:"createAt"` // 创建时间
	UpdateAt *time.Time `json:"updateAt"` // 更新时间
}
