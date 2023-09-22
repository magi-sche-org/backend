//go:generate mockgen -source=./user.go -destination=./mock/user.go -package=mockcontroller
package controller

import (
	"fmt"
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/appcontext"
	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/controller/internal/types"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	// Register(c echo.Context) error
	// GetEvents(c echo.Context) error
	Get(c echo.Context) error
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

func (uc *userController) Get(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := appcontext.Extract(ctx).GetUser()
	if err != nil {
		return err
	}

	ops, ouis, err := uc.uu.RetrieveUserProviders(ctx, user)
	if err != nil {
		return apperror.NewInternalError(err, "failed to retrieve user providers", "failed to retrieve user providers")
	}
	ur := types.UserResponse{
		User:      user,
		Providers: types.NewProviderResponse(ops, ouis),
	}
	return util.JSONResponse(c, http.StatusOK, ur)
}

// // Register implements UserController.
// func (uc *userController) Register(c echo.Context) error {
// 	user := entity.User{}
// 	if err := c.Bind(&user); err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	res, err := uc.uu.Register(c.Request().Context(), user)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, err.Error())
// 	}
// 	return util.JSONResponse(c, http.StatusOK, res)
// 	// return c.JSON(http.StatusOK, res)
// }

// GetEvents implements UserController.
// func (*userController) GetEvents(c echo.Context) error {
// 	panic("unimplemented")
// }

// GetExternalCalendars implements UserController.
func (uc *userController) GetExternalCalendars(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := appcontext.Extract(ctx).GetUser()
	if err != nil {
		return err
	}

	req := types.ExternalEventRequest{}
	if err := c.Bind(&req); err != nil {
		return apperror.NewInvalidRequestQueryError(err, nil)
	}
	fmt.Printf("req: %+v\n", req)

	events, err := uc.uu.FetchExternalCalendars(ctx, user, req.TimeMin, req.TimeMax)
	if err != nil {
		return apperror.NewInternalError(err, "failed to fetch external calendars", "failed to fetch external calendars")
	}

	return util.JSONResponse(c, http.StatusOK, events)
}
