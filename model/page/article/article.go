/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 17:03:00
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-11-13 21:51:05
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
	dbArticle := db.NewArticle()

	dbArticle.AID = article.AID
	dbArticle.Title = article.Title
	dbArticle.Content = article.Content
	dbArticle.AuthorID = article.AuthorID

	dbResult, err := dbArticle.AddArticle(ctx)

	if err != nil {
		msg := fmt.Errorf("SERVICE_ARTICLE_ARTICLE_AddArticle_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	res := &Article{
		ID:       dbResult.ID,
		AID:      dbResult.AID,
		Title:    dbResult.Title,
		Content:  dbResult.Content,
		AuthorID: dbResult.AuthorID,
		CreateAt: dbResult.CreatedAt,
		UpdateAt: dbResult.UpdatedAt,
	}

	return res, nil
}

func (article *Article) UpdateArticleByAID(ctx context.Context, aid uint64, params *UpdateArticleParams) error {
	dbArticle := db.NewArticle()

	val := map[string]interface{}{}

	if params.Title != nil {
		val[db.ArticleColumn.Title] = *params.Title
	}

	if params.Content != nil {
		val[db.ArticleColumn.Content] = *params.Content
	}

	_, err := dbArticle.UpdateArticleAID(ctx, aid, val)

	if err != nil {
		msg := fmt.Errorf("SERVICE_ARTICLE_ARTICLE_UpdateArticleByAID_DbArticleUpdateArticleAID_Failed")
		log.Println("err", msg)
		return msg
	}
	return nil
}

func (article *Article) DeleteArticleByAID(ctx context.Context, aid uint64) error {
	dbArticle := db.NewArticle()

	_, err := dbArticle.DeleteArticleByAID(ctx, aid)

	if err != nil {
		msg := fmt.Errorf("SERVICE_ARTICLE_ARTICLE_DeleteArticleByAID_DbArticleDeleteArticleByAID_Failed")
		log.Println("err", msg)
		return msg
	}

	return nil
}

func (article *Article) FindArticleByAID(ctx context.Context, aid uint64) (*Article, error) {
	dbArticle := db.NewArticle()

	dbResult, err := dbArticle.FindArticleByAID(ctx, aid)

	if err != nil {
		msg := fmt.Errorf("SERVICE_ARTICLE_ARTICLE_FindArticleByAID_DbArticleFindArticleByAID_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	result := &Article{
		ID:       dbResult.ID,
		AID:      dbArticle.AID,
		Title:    dbArticle.Title,
		Content:  dbArticle.Content,
		AuthorID: dbArticle.AuthorID,
		CreateAt: dbArticle.CreatedAt,
		UpdateAt: dbArticle.UpdatedAt,
	}
	return result, nil

}

type QueryArticleParam struct {
	Title    *string
	AuthorID *uint64
	Context  *string
}

func (article *Article) FindArticleList(ctx context.Context, pn int, num int, cond QueryArticleParam) (*ArticleList, error) {
	dbArticle := db.NewArticle()

	queryParams := make(map[string]interface{})

	if cond.Title != nil {
		queryParams[db.ArticleColumn.Title] = *cond.Title
	}

	if cond.AuthorID != nil {
		queryParams[db.ArticleColumn.AuthorID] = *cond.AuthorID
	}

	if cond.Context != nil {
		queryParams[db.ArticleColumn.Content] = *cond.Context
	}

	queryParams[db.ArticleColumn.IsDel] = 0

	dbResult, err := dbArticle.FindArticleList(ctx, pn, num, queryParams)

	if err != nil {
		msg := fmt.Errorf("SERVICE_ARTICLE_ARTICLE_FindArticleList_DbArticleFindArticleList_Failed")
		log.Println("err", msg)
		return nil, msg
	}

	result := make(ArticleList, 0)

	for _, item := range *dbResult {
		temp := Article{
			ID:       item.ID,
			AID:      item.AID,
			Title:    item.Title,
			Content:  item.Content,
			AuthorID: item.AuthorID,
			CreateAt: item.CreatedAt,
			UpdateAt: item.UpdatedAt,
		}

		result = append(result, temp)
	}

	return &result, nil
}
