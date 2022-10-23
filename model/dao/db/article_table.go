/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-22 15:59:41
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-10-23 17:58:34
 */
package db

import "time"

type article struct {
	ID        uint64     `gorm:"column:id;primaryKey;autoIncrement;" json:"id"`
	AID       uint64     `gorm:"column:aid;not null;comment:'文章id';" json:"aid"`
	Title     string     `gorm:"column:title;not null;comment:'标题';" json:"title"`
	Content   string     `gorm:"column:content;not null;comment:'内容';" json:"content"`
	AuthorID  uint64     `gorm:"column:author_id;not null;comment:'作者id';" json:"author_id"`
	CreatedAt *time.Time `gorm:"column:create_at;comment:'创建时间';not null;" json:"create_at"`
	UpdatedAt *time.Time `gorm:"column:update_at;comment:'更新时间';not null;" json:"update_at"`
	IsDel     uint8      `gorm:"column:is_del;comment:'已删除 0否 1是';not null; default:0;" json:"is_del"`
	Version   uint64     `gorm:"column:version; comment:'乐观锁';not null; default:0;" json:"version"`
}

type articleList []article

var ArticleColumn = struct {
	ID        string
	AID       string
	Title     string
	Content   string
	AuthorID  string
	CreatedAt string
	UpdatedAt string
	IsDel     string
	Version   string
}{
	ID:        "id",
	AID:       "aid",
	Title:     "title",
	Content:   "content",
	AuthorID:  "author_id",
	CreatedAt: "create_at",
	UpdatedAt: "update_at",
	IsDel:     "is_del",
	Version:   "version",
}

func (article) TableName() string {
	return "article"
}
