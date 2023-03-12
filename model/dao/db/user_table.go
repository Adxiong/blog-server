/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-22 15:59:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 17:13:10
 */
package db

import (
	"time"
)

// User [...]
type User struct {
	ID       uint64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	UID      uint64     `gorm:"index:idx_user_uid;column:uid;type:bigint unsigned;not null;comment:''用户唯一id''" json:"uid"`     // '用户唯一id'
	Password string     `gorm:"column:password;type:varchar(256);not null;comment:''密码''" json:"password"`                     // '密码'
	Username string     `gorm:"index:idx_username;column:username;type:varchar(256);not null;comment:''用户名''" json:"username"` // '用户名'
	Email    string     `gorm:"index:idx_email;column:email;type:varchar(256);not null;comment:''邮箱''" json:"email"`           // '邮箱'
	CreateAt *time.Time `gorm:"column:create_at;type:datetime;not null;comment:''创建时间''" json:"createAt"`                      // '创建时间'
	UpdateAt *time.Time `gorm:"column:update_at;type:datetime;not null;comment:''更新时间''" json:"updateAt"`                      // '更新时间'
	IsDel    uint8      `gorm:"column:is_del;type:tinyint unsigned;not null;default:0;comment:''已删除 0否 1是''" json:"isDel"`     // '已删除 0否 1是'
	Version  uint64     `gorm:"column:version;type:bigint unsigned;not null;default:0;comment:''乐观锁''" json:"version"`         // '乐观锁'
}

type UserList []User

// TableName get sql table name.获取数据库表名
func (m *User) TableName() string {
	return "user"
}

// UserColumns get sql column name.获取数据库列名
var UserColumns = struct {
	ID       string
	UID      string
	Password string
	Username string
	Email    string
	CreateAt string
	UpdateAt string
	IsDel    string
	Version  string
}{
	ID:       "id",
	UID:      "uid",
	Password: "password",
	Username: "username",
	Email:    "email",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	IsDel:    "is_del",
	Version:  "version",
}
