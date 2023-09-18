package converter

import (
	"context"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/models"
	"github.com/geekcamp-vol11-team30/backend/util"
)

func EventModelToEntity(ctx context.Context, em *models.Event, units []entity.EventTimeUnit, userAnswers []entity.UserEventAnswer) (entity.Event, error) {
	id, err := util.ULIDFromString(em.ID)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed to parse ULID(%s): %w", em.ID, err)
	}
	ownerId, err := util.ULIDFromString(em.OwnerID)
	if err != nil {
		return entity.Event{}, fmt.Errorf("failed to parse ULID(%s): %w", em.OwnerID, err)
	}
	return entity.Event{
		ID:            id,
		OwnerID:       ownerId,
		Name:          em.Name,
		Description:   em.Description,
		DurationAbout: em.DurationAbout,
		UnitSeconds:   int(em.UnitSeconds),
		Units:         units,
		UserAnswers:   userAnswers,
	}, nil
}
