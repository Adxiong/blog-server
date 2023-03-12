/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-22 15:59:41
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 17:10:33
 */
package db

import "time"

// Article [...]
type Article struct {
	ID       uint64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint unsigned;not null" json:"-"`
	Aid      uint64     `gorm:"column:aid;type:bigint unsigned;not null;comment:''文章id''" json:"aid"`                      // '文章id'
	AuthorID uint64     `gorm:"column:author_id;type:bigint unsigned;not null;comment:''作者id''" json:"authorId"`           // '作者id'
	Title    string     `gorm:"column:title;type:varchar(256);not null;comment:''标题''" json:"title"`                       // '标题'
	Tag      string     `gorm:"column:tag;type:varchar(256);not null;default:0;comment:''标签''" json:"tag"`                 // '标签'
	Content  string     `gorm:"column:content;type:varchar(256);not null;comment:''内容''" json:"content"`                   // '内容'
	Like     uint64     `gorm:"column:like;type:bigint unsigned;not null;default:0;comment:''点赞量''" json:"like"`           // '点赞量'
	CreateAt *time.Time `gorm:"column:create_at;type:datetime;not null;comment:''创建时间''" json:"createAt"`                  // '创建时间'
	UpdateAt *time.Time `gorm:"column:update_at;type:datetime;not null;comment:''更新时间''" json:"updateAt"`                  // '更新时间'
	IsDel    uint8      `gorm:"column:is_del;type:tinyint unsigned;not null;default:0;comment:''已删除 0否 1是''" json:"isDel"` // '已删除 0否 1是'
	Version  uint64     `gorm:"column:version;type:bigint unsigned;not null;default:0;comment:''乐观锁''" json:"version"`     // '乐观锁'
	ViewNum  uint64     `gorm:"column:view_num;type:bigint unsigned;not null;default:0;comment:''浏览量''" json:"viewNum"`    // '浏览量'
}

type ArticleList []Article

// TableName get sql table name.获取数据库表名
func (m *Article) TableName() string {
	return "article"
}

// ArticleColumns get sql column name.获取数据库列名
var ArticleColumns = struct {
	ID       string
	Aid      string
	AuthorID string
	Title    string
	Tag      string
	Content  string
	Like     string
	CreateAt string
	UpdateAt string
	IsDel    string
	Version  string
	ViewNum  string
}{
	ID:       "id",
	Aid:      "aid",
	AuthorID: "author_id",
	Title:    "title",
	Tag:      "tag",
	Content:  "content",
	Like:     "like",
	CreateAt: "create_at",
	UpdateAt: "update_at",
	IsDel:    "is_del",
	Version:  "version",
	ViewNum:  "view_num",
}
