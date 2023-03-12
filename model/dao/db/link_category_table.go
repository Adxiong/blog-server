/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2023-03-12 16:57:23
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 22:52:18
 */
package db

// LinkCategory [...]
type LinkCategory struct {
	ID         uint   `gorm:"autoIncrement:true;primaryKey;column:id;type:int unsigned;not null;comment:'自增id'" json:"-"` // 自增id
	Cid        uint   `gorm:"column:cid;type:int unsigned;not null;comment:'业务名称id'" json:"cid"`                          // 业务名称id
	Name       string `gorm:"column:name;type:varchar(255);not null;comment:'分类名称'" json:"name"`                          // 分类名称
	Version    uint   `gorm:"column:version;type:int unsigned;not null;default:0;comment:'乐观锁'" json:"version"`           // 乐观锁
	IsDel      uint   `gorm:"column:is_del;type:int unsigned;not null;default:0;comment:'删除标志位'" json:"isDel"`            // 删除标志位
	CreateTime uint   `gorm:"column:create_time;type:int unsigned;not null;comment:'创建时间'" json:"createTime"`             // 创建时间
	UpdateTime uint   `gorm:"column:update_time;type:int unsigned;not null;comment:'更新时间'" json:"updateTime"`             // 更新时间
}

type LinkCategoryList []LinkCategory

// TableName get sql table name.获取数据库表名
func (m *LinkCategory) TableName() string {
	return "link_category"
}

// LinkCategoryColumns get sql column name.获取数据库列名
var LinkCategoryColumns = struct {
	ID         string
	Cid        string
	Name       string
	Version    string
	IsDel      string
	CreateTime string
	UpdateTime string
}{
	ID:         "id",
	Cid:        "cid",
	Name:       "name",
	Version:    "version",
	IsDel:      "is_del",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}
