package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type queryArticleParam struct {
	Title    *string
	AuthorID *uint64
	Context  *string
}

func NewArticle() *Article {
	return &Article{}
}

func NewArticleList() *ArticleList {
	return &ArticleList{}
}

func (article *Article) AddArticle(ctx context.Context) (*Article, error) {
	err := GlobalDb.Table(article.TableName()).Create(article).Error
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return article, nil
}

func (Article *Article) UpdateArticleAID(ctx context.Context, aid uint64, values map[string]interface{}) (*int64, error) {
	res := GlobalDb.Table(Article.TableName()).Where("aid = ?", aid).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}
	return &res.RowsAffected, nil
}

func (Article *Article) DeleteArticleByAID(ctx context.Context, aid uint64) (*int64, error) {
	res := GlobalDb.Table(Article.TableName()).Where("aid = ?", aid).Update(ArticleColumns.IsDel, 1)
	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}
	return &res.RowsAffected, nil
}

func (Article *Article) FindArticleByAID(ctx context.Context, aid uint64) (*Article, error) {
	params := map[string]interface{}{
		ArticleColumns.Aid:   aid,
		ArticleColumns.IsDel: 0,
	}
	res := GlobalDb.Table(Article.TableName()).Where(params).Limit(1).Find(Article)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return Article, nil
}

func (Article *Article) FindArticleList(ctx context.Context, pn int, num int, params map[string]interface{}) (*ArticleList, error) {
	ArticleList := NewArticleList()
	offset := pn * num

	res := GlobalDb.Table(Article.TableName()).Offset(offset).Limit(num).Where(params).Find(ArticleList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return ArticleList, nil
}

func (Article *Article) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	Article.CreateAt = &t
	Article.UpdateAt = &t

	return nil
}

func (Article *Article) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[ArticleColumns.UpdateAt]; !ok {
			t := time.Now()
			values[ArticleColumns.UpdateAt] = &t
		}

		if _, ok := values[ArticleColumns.Version]; !ok {
			values[ArticleColumns.Version] = gorm.Expr(fmt.Sprintf("%s + ?", ArticleColumns.Version), 1)
		}
	}

	return nil
}
