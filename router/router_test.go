package router_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geekcamp-vol11-team30/backend/config"
	mockcontroller "github.com/geekcamp-vol11-team30/backend/controller/mock"
	mockmiddleware "github.com/geekcamp-vol11-team30/backend/middleware/mock"
	"github.com/geekcamp-vol11-team30/backend/router"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestNewRouter(t *testing.T) {
	logger := zap.NewExample()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	uc := mockcontroller.NewMockUserController(ctrl)
	ac := mockcontroller.NewMockAuthController(ctrl)
	ec := mockcontroller.NewMockEventController(ctrl)

	em := mockmiddleware.NewMockErrorMiddleware(ctrl)
	atm := mockmiddleware.NewMockAccessTimeMiddleware(ctrl)
	am := mockmiddleware.NewMockAuthMiddleware(ctrl)
	e := router.NewRouter(&config.Config{}, logger, em, atm, am, uc, ac, ec)

	t.Run("GET /health", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/health", nil)
		recorder := httptest.NewRecorder()

		context := e.NewContext(request, recorder)

		e.Router().Find(echo.GET, "/health", context)

		if assert.NoError(t, context.Handler()(context)) {
			assert.Equal(t, http.StatusOK, recorder.Code)
			assert.Equal(t, "OK", recorder.Body.String())
		}
	})
}
