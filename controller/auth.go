//go:generate mockgen -source=./auth.go -destination=./mock/auth.go -package=mockcontroller
package controller

import (
	"net/http"

	"github.com/geekcamp-vol11-team30/backend/config"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/usecase"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/labstack/echo/v4"
)

type AuthController interface {
	// 未登録ユーザー作成・トークン発行
	CreateUnregisteredUserAndToken(c echo.Context) error
	// トークン更新
	RefreshToken(c echo.Context) error
	// CSRFトークン発行
	CreateCSRFToken(c echo.Context) error
}

type authController struct {
	cfg *config.Config
	uu  usecase.UserUsecase
	au  usecase.AuthUsecase
}

func NewAuthController(cfg *config.Config, uu usecase.UserUsecase, au usecase.AuthUsecase) AuthController {
	return &authController{
		cfg: cfg,
		uu:  uu,
		au:  au,
	}
}

// Slash implements SlackController.
func (ac *authController) CreateUnregisteredUserAndToken(c echo.Context) error {
	ctx := c.Request().Context()
	user, err := ac.uu.CreateAnonymousUser(ctx)
	if err != nil {
		return err
	}
	token, err := ac.au.CreateToken(ctx, user)
	if err != nil {
		return err
	}
	util.SetTokenCookie(c, *ac.cfg, token)
	return util.JSONResponse(c, http.StatusOK, token)
	// return c.JSON(200, token)
	// panic("unimplemented")
}

// RefreshToken implements AuthController.
func (ac *authController) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()
	rtokenReq := entity.RefreshTokenRequest{}
	if err := c.Bind(&rtokenReq); err != nil {
		return err
	}
	token, err := ac.au.RefreshToken(ctx, rtokenReq.RefreshToken)
	if err != nil {
		return err
	}
	// http.StatusNotFound
	return util.JSONResponse(c, http.StatusOK, token)
	// return c.JSON(200, token)
}

// CreateCSRFToken implements AuthController.
func (*authController) CreateCSRFToken(c echo.Context) error {
	token, ok := c.Get("csrf").(string)
	if !ok {
		return util.JSONResponse(c, http.StatusOK, map[string]string{
			"csrf": "",
		})
	}
	return util.JSONResponse(c, http.StatusOK, map[string]string{
		"csrf": token,
	})
}
