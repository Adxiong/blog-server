/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 17:03:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-03-12 17:13:56
 */
package svrarticle

import (
	"blogserver/model/dao/db"
	"context"
	"fmt"
	"log"
	"time"
)

type ArticleService interface {
	AddArticle()
	UpdateArticleByUID()
	DeleteArticleByUID()
	FindArticleByUID()
	FindArticleList()
}

type Article struct {
	ID       uint64     `json:"id"`
	AID      uint64     `json:"aid"`
	Title    string     `json:"title"`
	Content  string     `json:"content"`
	AuthorID uint64     `json:"author_id"`
	CreateAt *time.Time `json:"create_at"`
	UpdateAt *time.Time `json:"update_at"`
}

type ArticleList []Article

type UpdateArticleParams struct {
	Title   *string
	Content *string
}

func NewArticleService() *Article {
	return &Article{}
}

func (article *Article) AddArticle(ctx context.Context) (*Article, error) {
	var result *Article
	var err error

	dbArticle := db.NewArticle()

	dbArticle.Aid = article.AID
	dbArticle.Title = article.Title
	dbArticle.Content = article.Content
	dbArticle.AuthorID = article.AuthorID

	dbResult, dbResultErr := dbArticle.AddArticle(ctx)

	if dbResultErr != nil {
		err = fmt.Errorf("SERVICE_ARTICLE_ARTICLE_AddArticle_Failed")
		log.Println("err", err)
		return result, err
	}

	result = &Article{
		ID:       dbResult.ID,
		AID:      dbResult.Aid,
		Title:    dbResult.Title,
		Content:  dbResult.Content,
		AuthorID: dbResult.AuthorID,
		CreateAt: dbResult.CreateAt,
		UpdateAt: dbResult.UpdateAt,
	}

	return result, err
}

func (article *Article) UpdateArticleByAID(ctx context.Context, aid uint64, params *UpdateArticleParams) error {
	var err error

	dbArticle := db.NewArticle()

	val := map[string]interface{}{}

	if params.Title != nil {
		val[db.ArticleColumns.Title] = *params.Title
	}

	if params.Content != nil {
		val[db.ArticleColumns.Content] = *params.Content
	}

	rowsAffected, UpdateErr := dbArticle.UpdateArticleAID(ctx, aid, val)

	if UpdateErr != nil || int(*rowsAffected) <= 0 {
		err = fmt.Errorf("SERVICE_ARTICLE_ARTICLE_UpdateArticleByAID_DbArticleUpdateArticleAID_Failed")
		log.Println("err", err)
		return err
	}

	return err
}

func (article *Article) DeleteArticleByAID(ctx context.Context, aid uint64) error {
	var err error

	dbArticle := db.NewArticle()

	rowsAffected, DelteErr := dbArticle.DeleteArticleByAID(ctx, aid)

	if DelteErr != nil || int(*rowsAffected) <= 0 {
		err = fmt.Errorf("SERVICE_ARTICLE_ARTICLE_DeleteArticleByAID_DbArticleDeleteArticleByAID_Failed")
		log.Println("err", err)
		return err
	}

	return err
}

func (article *Article) FindArticleByAID(ctx context.Context, aid uint64) (*Article, error) {
	var result *Article
	var err error

	dbArticle := db.NewArticle()

	dbResult, dbResultErr := dbArticle.FindArticleByAID(ctx, aid)

	if dbResultErr != nil {
		err = fmt.Errorf("SERVICE_ARTICLE_ARTICLE_FindArticleByAID_DbArticleFindArticleByAID_Failed")
		return result, err
	}

	if dbResult.Aid <= 0 {
		return result, err
	}

	result = &Article{
		ID:       dbResult.ID,
		AID:      dbArticle.Aid,
		Title:    dbArticle.Title,
		Content:  dbArticle.Content,
		AuthorID: dbArticle.AuthorID,
		CreateAt: dbArticle.CreateAt,
		UpdateAt: dbArticle.UpdateAt,
	}
	return result, err

}

type QueryArticleParam struct {
	Title    *string
	AuthorID *uint64
	Context  *string
}

func (article *Article) FindArticleList(ctx context.Context, pn int, num int, cond QueryArticleParam) (*ArticleList, error) {
	var result ArticleList
	var err error

	dbArticle := db.NewArticle()

	queryParams := make(map[string]interface{})

	if cond.Title != nil {
		queryParams[db.ArticleColumns.Title] = *cond.Title
	}

	if cond.AuthorID != nil {
		queryParams[db.ArticleColumns.AuthorID] = *cond.AuthorID
	}

	if cond.Context != nil {
		queryParams[db.ArticleColumns.Content] = *cond.Context
	}

	queryParams[db.ArticleColumns.IsDel] = 0

	dbResult, dbResultErr := dbArticle.FindArticleList(ctx, pn, num, queryParams)

	if dbResultErr != nil {
		err = fmt.Errorf("SERVICE_ARTICLE_ARTICLE_FindArticleList_DbArticleFindArticleList_Failed")
		log.Println("err", err)
		return &result, err
	}

	if len(*dbResult) <= 0 {
		return &result, err
	}

	result = make(ArticleList, 0)

	for _, item := range *dbResult {
		temp := Article{
			ID:       item.ID,
			AID:      item.Aid,
			Title:    item.Title,
			Content:  item.Content,
			AuthorID: item.AuthorID,
			CreateAt: item.CreateAt,
			UpdateAt: item.UpdateAt,
		}

		result = append(result, temp)
	}

	return &result, nil
}
