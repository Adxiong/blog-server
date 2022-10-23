/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-22 15:59:29
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 15:39:30
 */
package db

import (
	"time"
)

type user struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	UID       uint64     `gorm:"column:uid;comment:'用户唯一id';not null;index;" json:"uid"`
	Password  string     `gorm:"column:password;comment:'密码';not null;" json:"password"`
	Username  string     `gorm:"column:username;comment:'用户名';not null;" json:"username"`
	Email     string     `gorm:"column:email;comment:'邮箱';not null;" json:"email"`
	CreatedAt *time.Time `gorm:"column:create_at;comment:'创建时间';not null;" json:"create_at"`
	UpdatedAt *time.Time `gorm:"column:update_at;comment:'更新时间';not null;" json:"update_at"`
	IsDel     uint8      `gorm:"column:is_del;comment:'已删除 0否 1是';not null; default:0;" json:"is_del"`
	Version   uint64     `gorm:"column:version; comment:'乐观锁';not null; default:0;" json:"version"`
}

type userList []user

var UserColumn = struct {
	ID        string
	UID       string
	Password  string
	Username  string
	Email     string
	CreatedAt string
	UpdatedAt string
	IsDel     string
	Version   string
}{
	ID:        "id",
	UID:       "uid",
	Password:  "password",
	Username:  "username",
	Email:     "email",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	IsDel:     "is_del",
	Version:   "version",
}

func (user) TableName() string {
	return "user"
}
