/*
 * @Description:

 * @version:
 * @Author: Adxiong
 * @Date: 2022-10-23 22:34:28
 * @LastEditors: Adxiong
 * @LastEditTime: 2023-05-02 16:11:33
 */
package httpserver

import (
	"blogserver/controller/article"
	"blogserver/controller/link"
	"blogserver/controller/user"
	"blogserver/library/request"
	"context"

	"github.com/gin-gonic/gin"
)

func RegisterController(ctx context.Context, r *gin.Engine) {
	for key, item := range routes {
		r.Handle(item.Method, key, item.Handle...)
	}
}

type RouterHandler struct {
	Method string
	Handle []gin.HandlerFunc
	Filter func()
}

type HandleFuncs []gin.HandlerFunc

var routes = map[string]RouterHandler{
	"/user/register": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.AddUser,
		},
	},
	"/user/del": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.DelUser,
		},
	},
	"/user/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			user.UpdateUser,
		},
	},
	"/user/list": {
		Method: request.METHODGET,
		Handle: HandleFuncs{
			user.GetUserList,
		},
	},
	"/article/detail": {
		Method: request.METHODGET,
		Handle: HandleFuncs{
			article.GetArticleDetail,
		},
	},
	"/article/add": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			article.AddArticle,
		},
	},
	"/article/update": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			article.UpdateArticle,
		},
	},
	"/article/del": {
		Method: request.METHODPOST,
		Handle: HandleFuncs{
			article.DelArticle,
		},
	},
	"/article/list": {
		Method: request.METHODGET,
		Handle: HandleFuncs{
			article.GetArticleList,
		},
	},
	"/category/list": {
		Method: request.METHODGET,
		Handle: HandleFuncs{
			link.GetLinkCategoryList,
		},
	},
	"/link/list": {
		Method: request.METHODGET,
		Handle: HandleFuncs{
			link.GetLinkList,
		},
	},
}
