/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:57:38
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 17:18:05
 */
package db

// Link [...]
type Link struct {
	ID         uint   `gorm:"primaryKey;column:id;type:int unsigned;not null;comment:'自增id'" json:"-"`          // 自增id
	Name       string `gorm:"column:name;type:varchar(255);not null;comment:'link名称'" json:"name"`              // link名称
	NameID     uint   `gorm:"column:name_id;type:int unsigned;not null;comment:'link_id'" json:"nameId"`        // link_id
	Pic        string `gorm:"column:pic;type:varchar(255);default:null;comment:'图片地址'" json:"pic"`              // 图片地址
	Brief      string `gorm:"column:brief;type:varchar(255);not null;comment:'简介'" json:"brief"`                // 简介
	Version    uint   `gorm:"column:version;type:int unsigned;not null;default:0;comment:'乐观锁'" json:"version"` // 乐观锁
	IsDel      uint   `gorm:"column:is_del;type:int unsigned;not null;default:0;comment:'删除标志位'" json:"isDel"`  // 删除标志位
	CreateTime uint   `gorm:"column:create_time;type:int unsigned;not null;comment:'创建时间'" json:"createTime"`   // 创建时间
	UpdateTime uint   `gorm:"column:update_time;type:int unsigned;not null;comment:'更新时间'" json:"updateTime"`   // 更新时间
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
	NameID     string
	Pic        string
	Brief      string
	Version    string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Name:       "name",
	NameID:     "name_id",
	Pic:        "pic",
	Brief:      "brief",
	Version:    "version",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
