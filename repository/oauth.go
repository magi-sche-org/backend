package repository

import (
	"context"
	"database/sql"

	"github.com/geekcamp-vol11-team30/backend/db/models"
	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type OauthRepository interface {
	RegisterProvider(ctx context.Context, op entity.OauthProvider) (entity.OauthProvider, error)
	FetchProviderByName(ctx context.Context, name string) (entity.OauthProvider, error)
	RegisterOauthUserInfo(ctx context.Context, oui entity.OauthUserInfo) (entity.OauthUserInfo, error)
	FetchOauthUserInfos(ctx context.Context, user entity.User) ([]entity.OauthUserInfo, error)
	FetchOauthUserInfo(ctx context.Context, providerId ulid.ULID, user entity.User) (entity.OauthUserInfo, error)
}

type oauthRepository struct {
	db *sql.DB
}

func NewOauthRepository(db *sql.DB) OauthRepository {
	return &oauthRepository{
		db: db,
	}
}

// RegisterProvider implements OauthRepository.
func (oar *oauthRepository) RegisterProvider(ctx context.Context, op entity.OauthProvider) (entity.OauthProvider, error) {
	id := util.GenerateULIDNow()
	opm := &models.OauthProvider{
		ID:           util.ULIDToString(id),
		Name:         op.Name,
		ClientID:     op.ClientId,
		ClientSecret: op.ClientSecret,
	}
	err := opm.Upsert(ctx, oar.db, boil.Infer(), boil.Infer())
	if err != nil {
		return entity.OauthProvider{}, err
	}
	return entity.OauthProvider{
		ID:           id,
		Name:         opm.Name,
		ClientId:     opm.ClientID,
		ClientSecret: opm.ClientSecret,
	}, nil
}

// FetchProviderByName implements OauthRepository.
func (or *oauthRepository) FetchProviderByName(ctx context.Context, name string) (entity.OauthProvider, error) {
	oapm, err := models.OauthProviders(models.OauthProviderWhere.Name.EQ(name)).One(ctx, or.db)
	if err != nil {
		return entity.OauthProvider{}, err
	}
	id, err := util.ULIDFromString(oapm.ID)
	if err != nil {
		return entity.OauthProvider{}, err
	}
	return entity.OauthProvider{
		ID:           id,
		Name:         oapm.Name,
		ClientId:     oapm.ClientID,
		ClientSecret: oapm.ClientSecret,
	}, nil
}

// RegisterOauthUserInfo implements OauthRepository.
func (oar *oauthRepository) RegisterOauthUserInfo(ctx context.Context, oui entity.OauthUserInfo) (entity.OauthUserInfo, error) {
	id := util.GenerateULID(ctx)
	ouim := &models.OauthUserInfo{
		ID:                    util.ULIDToString(id),
		UserID:                util.ULIDToString(oui.UserId),
		ProviderID:            util.ULIDToString(oui.ProviderId),
		ProviderUID:           oui.ProviderUid,
		AccessToken:           oui.AccessToken,
		RefreshToken:          oui.RefreshToken,
		AccessTokenExpiresAt:  oui.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: null.TimeFromPtr(oui.RefreshTokenExpiresAt),
	}
	err := ouim.Upsert(ctx, oar.db, boil.Infer(), boil.Infer())
	if err != nil {
		return entity.OauthUserInfo{}, err
	}
	userId, _ := util.ULIDFromString(ouim.UserID)
	providerId, _ := util.ULIDFromString(ouim.ProviderID)
	return entity.OauthUserInfo{
		ID:                    id,
		UserId:                userId,
		ProviderId:            providerId,
		ProviderUid:           ouim.ProviderUID,
		AccessToken:           ouim.AccessToken,
		RefreshToken:          ouim.RefreshToken,
		AccessTokenExpiresAt:  ouim.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: ouim.RefreshTokenExpiresAt.Ptr(),
	}, nil
}

// FetchOauthUserInfo implements OauthRepository.
func (oar *oauthRepository) FetchOauthUserInfos(ctx context.Context, user entity.User) ([]entity.OauthUserInfo, error) {
	ouism, err := models.OauthUserInfos(
		models.OauthUserInfoWhere.UserID.EQ(util.ULIDToString(user.ID)),
	).All(ctx, oar.db)
	if err != nil {
		return nil, err
	}
	ouis := make([]entity.OauthUserInfo, len(ouism))
	for i, oui := range ouism {
		id, _ := util.ULIDFromString(oui.ID)
		userId, _ := util.ULIDFromString(oui.UserID)
		providerId, _ := util.ULIDFromString(oui.ProviderID)
		ouis[i] = entity.OauthUserInfo{
			ID:                    id,
			UserId:                userId,
			ProviderId:            providerId,
			ProviderUid:           oui.ProviderUID,
			AccessToken:           oui.AccessToken,
			RefreshToken:          oui.RefreshToken,
			AccessTokenExpiresAt:  oui.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: oui.RefreshTokenExpiresAt.Ptr(),
		}
	}
	return ouis, nil
}

// FetchOauthUserInfo implements OauthRepository.
func (oar *oauthRepository) FetchOauthUserInfo(ctx context.Context, providerId ulid.ULID, user entity.User) (entity.OauthUserInfo, error) {
	ouim, err := models.OauthUserInfos(
		models.OauthUserInfoWhere.ProviderUID.EQ(util.ULIDToString(providerId)),
		models.OauthUserInfoWhere.UserID.EQ(util.ULIDToString(user.ID)),
	).One(ctx, oar.db)
	if err != nil {
		return entity.OauthUserInfo{}, err
	}
	id, _ := util.ULIDFromString(ouim.ID)
	userId, _ := util.ULIDFromString(ouim.UserID)
	// providerId, _ := util.ULIDFromString(ouim.ProviderID)
	return entity.OauthUserInfo{
		ID:                    id,
		UserId:                userId,
		ProviderId:            providerId,
		ProviderUid:           ouim.ProviderUID,
		AccessToken:           ouim.AccessToken,
		RefreshToken:          ouim.RefreshToken,
		AccessTokenExpiresAt:  ouim.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: ouim.RefreshTokenExpiresAt.Ptr(),
	}, nil
}
