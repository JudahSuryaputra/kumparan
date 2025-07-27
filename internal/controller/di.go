package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"kumparan/internal/controller/article"
	"kumparan/internal/controller/middleware"
	"kumparan/internal/shared"
)

type (
	Holder struct {
		dig.In
		Deps               shared.Deps
		InternalMiddleware middleware.Middleware
		PlatformController *PlatformController
		ArticleController  *article.Controller
	}
)

func Register(container *dig.Container) error {
	if err := middleware.Register(container); err != nil {
		return err
	}
	if err := container.Provide(NewPlatformController); err != nil {
		return err
	}
	if err := container.Provide(article.NewArticleController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {
	// check app health
	app.GET("/health", h.PlatformController.CheckSelf)

	v1 := app.Group("/v1")
	{
		v1.POST("/article", h.ArticleController.Post)
		v1.GET("/articles", h.ArticleController.GetArticles)
	}

}
