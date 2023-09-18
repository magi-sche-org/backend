package converter

import (
	"context"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/models"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/volatiletech/null/v8"
)

func OauthUserInfoModelToEntity(ctx context.Context, ouim *models.OauthUserInfo) (entity.OauthUserInfo, error) {
	id, err := util.ULIDFromString(ouim.ID)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to parse ULID(%s): %w", ouim.ID, err)
	}
	userId, err := util.ULIDFromString(ouim.UserID)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to parse ULID(%s): %w", ouim.UserID, err)
	}
	providerId, err := util.ULIDFromString(ouim.ProviderID)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to parse ULID(%s): %w", ouim.ProviderID, err)
	}

	var provider *entity.OauthProvider
	provider = nil
	if pr := ouim.R; pr != nil {
		if p := pr.Provider; p != nil {
			provider, err = OauthProviderModelToEntity(ctx, p)
			if err != nil {
				return entity.OauthUserInfo{}, fmt.Errorf("failed to convert OauthProviderModel to entity on OauthUserInfoModelToEntity: %w", err)
			}
		}
	}

	return entity.OauthUserInfo{
		ID:                    id,
		UserId:                userId,
		ProviderId:            providerId,
		ProviderUid:           ouim.ProviderUID,
		AccessToken:           ouim.AccessToken,
		RefreshToken:          ouim.RefreshToken,
		AccessTokenExpiresAt:  ouim.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: ouim.RefreshTokenExpiresAt.Ptr(),
		Provider:              provider,
	}, nil
}

func OauthUserInfoSliceModelToEntity(ctx context.Context, ouims models.OauthUserInfoSlice) ([]entity.OauthUserInfo, error) {
	ouis := make([]entity.OauthUserInfo, len(ouims))
	for i, oui := range ouims {
		ouii, err := OauthUserInfoModelToEntity(ctx, oui)
		if err != nil {
			return nil, fmt.Errorf("failed to convert OauthUserInfoModel to entity on OauthUserInfoSliceModelToEntity: %w", err)
		}
		ouis[i] = ouii
	}
	return ouis, nil
}

func OauthUserInfoEntityToModel(ctx context.Context, oui entity.OauthUserInfo) *models.OauthUserInfo {
	return &models.OauthUserInfo{
		ID:                    util.ULIDToString(oui.ID),
		UserID:                util.ULIDToString(oui.UserId),
		ProviderID:            util.ULIDToString(oui.ProviderId),
		ProviderUID:           oui.ProviderUid,
		AccessToken:           oui.AccessToken,
		RefreshToken:          oui.RefreshToken,
		AccessTokenExpiresAt:  oui.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: null.TimeFromPtr(oui.RefreshTokenExpiresAt),
	}
}

func OauthProviderModelToEntity(ctx context.Context, opm *models.OauthProvider) (*entity.OauthProvider, error) {
	id, err := util.ULIDFromString(opm.ID)
	if err != nil {
		return &entity.OauthProvider{}, fmt.Errorf("failed to parse ULID(%s): %w", opm.ID, err)
	}
	return &entity.OauthProvider{
		ID:           id,
		Name:         opm.Name,
		ClientId:     opm.ClientID,
		ClientSecret: opm.ClientSecret,
	}, nil
}

func OauthProviderEntityToModel(ctx context.Context, op entity.OauthProvider) *models.OauthProvider {
	return &models.OauthProvider{
		ID:           util.ULIDToString(op.ID),
		Name:         op.Name,
		ClientID:     op.ClientId,
		ClientSecret: op.ClientSecret,
	}
}
