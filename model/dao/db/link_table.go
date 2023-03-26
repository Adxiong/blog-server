/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:57:38
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-26 23:13:20
 */
package db

import "time"

// Link [...]
type Link struct {
	ID         uint64     `gorm:"primaryKey;column:id;type:int unsigned;not null;comment:'自增id'" json:"id"`            // 自增id
	Name       string     `gorm:"column:name;type:varchar(255);not null;comment:'link名称'" json:"name"`                 // link名称
	Url        string     `gorm:"column:url;type:varchar(255);not null;comment:'link地址'" json:"url"`                   // link地址
	LinkID     uint64     `gorm:"column:link_id;type:int unsigned;not null;comment:'link_id'" json:"link_id"`          // link_id
	CategoryID uint64     `gorm:"column:category_id;type:int unsigned;not null;comment:'link类别id'" json:"category_id"` // link类别id
	Pic        string     `gorm:"column:pic;type:varchar(255);default:null;comment:'图片地址'" json:"pic"`                 // 图片地址
	Brief      string     `gorm:"column:brief;type:varchar(255);not null;comment:'简介'" json:"brief"`                   // 简介
	Version    uint       `gorm:"column:version;type:int unsigned;not null;default:0;comment:'乐观锁'" json:"version"`    // 乐观锁
	IsDel      uint       `gorm:"column:is_del;type:int unsigned;not null;default:0;comment:'删除标志位'" json:"is_del"`    // 删除标志位
	CreateAt   *time.Time `gorm:"column:create_at;type:datetime;not null;comment:'创建时间'" json:"create_at"`             // 创建时间
	UpdateAt   *time.Time `gorm:"column:update_at;type:datetime;not null;comment:'更新时间'" json:"update_at"`             // 更新时间
}
type LinkList []Link

// TableName get sql table name.获取数据库表名
func (m *Link) TableName() string {
	return "link"
}

// LinkColumns get sql column name.获取数据库列名
var LinkColumns = struct {
	ID         string
	Name       string
	Url        string
	LinkID     string
	CategoryID string
	Pic        string
	Brief      string
	Version    string
	IsDel      string
	CreateAt   string
	UpdateAt   string
}{
	ID:         "id",
	Name:       "name",
	Url:        "url",
	LinkID:     "link_id",
	CategoryID: "category_id",
	Pic:        "pic",
	Brief:      "brief",
	Version:    "version",
	IsDel:      "is_del",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}
