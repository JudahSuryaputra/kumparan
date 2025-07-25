package controller

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
	"kumparan/internal/controller/middleware"
	"kumparan/internal/controller/user"
	"kumparan/internal/shared"
)

type (
	Holder struct {
		dig.In
		Deps               shared.Deps
		InternalMiddleware middleware.Middleware
		PlatformController *PlatformController
		UserController     *user.Controller
	}
)

func Register(container *dig.Container) error {
	if err := middleware.Register(container); err != nil {
		return err
	}
	if err := container.Provide(NewPlatformController); err != nil {
		return err
	}
	if err := container.Provide(user.NewUserController); err != nil {
		return err
	}

	return nil
}

func (h *Holder) SetupRoutes(app *echo.Echo) {
	// check app health
	app.GET("/health", h.PlatformController.CheckSelf)

	v1 := app.Group("/v1")
	{
		v1.GET("/user", h.UserController.GetUserByID)
	}

}
