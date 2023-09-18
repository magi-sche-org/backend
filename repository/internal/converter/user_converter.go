package converter

import (
	"context"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/models"
	"github.com/geekcamp-vol11-team30/backend/util"
)

func UserModelToEntity(ctx context.Context, um *models.User) (entity.User, error) {
	id, err := util.ULIDFromString(um.ID)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to parse ULID(%s): %w", um.ID, err)
	}
	return entity.User{
		ID:           id,
		Name:         um.Name,
		IsRegistered: um.IsRegistered,
	}, nil
}

func UserEntityToModel(ctx context.Context, ue entity.User) *models.User {
	return &models.User{
		ID:           util.ULIDToString(ue.ID),
		Name:         ue.Name,
		IsRegistered: ue.IsRegistered,
	}
}
