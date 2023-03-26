package link

import (
	"blogserver/library/response"
	"blogserver/model/page/link"
	"fmt"

	"github.com/gin-gonic/gin"
)

type AddLinkCategoryParams struct {
	Name string `json:"name" binding:"required"`
}

type AddLinkCategoryResponse struct {
	ID uint64 `json:"id"`
}

func AddLinkCategory(ctx *gin.Context) {
	params, errParams := parseAddLinkCategoryParams(ctx)
	if errParams != nil {
		err := fmt.Errorf("params is not valid")
		fmt.Println("err", errParams)
		response.Error(ctx, 200, err.Error())
	}

	svrLcg := link.LinkCategory{}
	Res, errRes := svrLcg.AddLinkCategory(ctx, params.Name)

	if errRes != nil {
		err := fmt.Errorf("添加失败")
		fmt.Println("err", errParams)
		response.Error(ctx, 200, err.Error())
	}

	result := &AddLinkCategoryResponse{
		ID: Res.ID,
	}

	response.Json(ctx, result)

}

func parseAddLinkCategoryParams(ctx *gin.Context) (*AddLinkCategoryParams, error) {
	p := &AddLinkCategoryParams{}

	err := ctx.ShouldBindJSON(p)

	if err != nil {
		return nil, err
	}

	return p, nil
}
