package appcontext

import (
	"context"
	"time"

	"github.com/geekcamp-vol11-team30/backend/apperror"
	"github.com/geekcamp-vol11-team30/backend/entity"
)

type contextKey struct{}

type AppContext struct {
	Now  time.Time
	user *entity.User
}

func Set(ctx context.Context, actx *AppContext) context.Context {
	return context.WithValue(ctx, contextKey{}, actx)
}

func Extract(ctx context.Context) *AppContext {
	value := ctx.Value(contextKey{})
	if value == nil {
		return nil
	}
	actx, ok := value.(*AppContext)
	if !ok {
		return nil
	}
	return actx
}

func (actx *AppContext) SetUser(user entity.User) {
	actx.user = &user
}

func (actx *AppContext) GetUser() (entity.User, error) {
	user := actx.user
	if user == nil {
		return entity.User{}, apperror.NewInternalError(nil, nil, "user is nil")
	}
	return *actx.user, nil
}
