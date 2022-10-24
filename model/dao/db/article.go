package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"
)

type ArticleDao interface {
	AddArticle()
	UpdateArticleAID()
	DeleteArticleByAID()
	FindArticleByAID()
	FindArticleList()
}

func NewArticle() *article {
	return &article{}
}

func NewArticleList() *articleList {
	return &articleList{}
}

func (article *article) AddArticle(ctx context.Context) (*article, error) {
	err := GlobalDb.Table(article.TableName()).Create(article).Error
	if err != nil {
		log.Println("err", err)
		return nil, err
	}
	return article, nil
}

func (article *article) UpdateArticleAID(ctx context.Context, aid uint64, values map[string]interface{}) (*int64, error) {
	res := GlobalDb.Table(article.TableName()).Where("aid = ?", aid).Updates(values)
	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}
	return &res.RowsAffected, nil
}

func (article *article) DeleteArticleByAID(ctx context.Context, aid uint64) (*int64, error) {
	res := GlobalDb.Table(article.TableName()).Where("aid = ?", aid).Update(ArticleColumn.IsDel, 1)
	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}
	return &res.RowsAffected, nil
}

func (article *article) FindArticleByAID(ctx context.Context, aid uint64) (*article, error) {
	params := map[string]interface{}{
		ArticleColumn.AID: aid,
	}
	res := GlobalDb.Table(article.TableName()).Where(params).Limit(1).Find(article)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return article, nil
}

func (article *article) FindArticleList(ctx context.Context, pn int, num int, params map[string]interface{}) (*articleList, error) {
	articleList := NewArticleList()
	offset := pn * num

	res := GlobalDb.Table(article.TableName()).Offset(offset).Limit(num).Where(params).Find(articleList)

	if res.Error != nil {
		log.Println("err", res.Error)
		return nil, res.Error
	}

	return articleList, nil
}

func (article *article) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	article.CreatedAt = &t
	article.UpdatedAt = &t

	return nil
}

func (article *article) BeforeUpdate(tx *gorm.DB) error {
	if values, ok := tx.Statement.Dest.(map[string]interface{}); ok {
		if _, ok := values[ArticleColumn.UpdatedAt]; !ok {
			t := time.Now()
			values[ArticleColumn.UpdatedAt] = &t
		}

		if _, ok := values[ArticleColumn.Version]; !ok {
			values[ArticleColumn.Version] = gorm.Expr(fmt.Sprintf("%s + ?", ArticleColumn.Version), 1)
		}
	}

	return nil
}
