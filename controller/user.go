//go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mockcontroller
package controller

import (
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Register(c echo.Context) error
	GetEvents(c echo.Context) error
	GetExternalCalendars(c echo.Context) error
}

type userController struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &userController{
		uu: uu,
	}
}

// Register implements UserController.
func (uc *userController) Register(c echo.Context) error {
	user := entity.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	res, err := uc.uu.Register(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return util.JSONResponse(c, http.StatusOK, res)
	// return c.JSON(http.StatusOK, res)
}

// GetEvents implements UserController.
func (*userController) GetEvents(c echo.Context) error {
	panic("unimplemented")
}

// GetExternalCalendars implements UserController.
func (uc *userController) GetExternalCalendars(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := appcontext.Extract(ctx).GetUser()
	if err != nil {
		return err
	}

}
