package article

import (
	"github.com/labstack/echo/v4"
	"kumparan/internal/controller/helper"
	"kumparan/internal/service"
	"kumparan/internal/shared"
	"kumparan/internal/shared/dto"
	"strconv"
)

type (
	Controller struct {
		deps     shared.Deps
		services service.Holder
	}
)

func NewArticleController(deps shared.Deps, services service.Holder) (*Controller, error) {
	return &Controller{deps: deps, services: services}, nil
}

func (c *Controller) GetArticles(ctx echo.Context) (err error) {
	var (
		rctx = ctx.Request().Context()
	)

	filter := dto.ArticleFilter{
		Query:  ctx.QueryParam("query"),
		Author: ctx.QueryParam("author"),
		Limit:  10,
		Offset: 0,
	}

	if limit := ctx.QueryParam("limit"); limit != "" {
		if v, err := strconv.Atoi(limit); err == nil && v > 0 {
			filter.Limit = v
		}
	}
	if offset := ctx.QueryParam("offset"); offset != "" {
		if v, err := strconv.Atoi(offset); err == nil && v >= 0 {
			filter.Offset = v
		}
	}

	resp, err := c.services.Article.Get(rctx, &filter)
	if err != nil {
		//return helper.FailWithDataResponse(ctx, err)
	}

	return helper.SuccessResponse(ctx, resp)
}
