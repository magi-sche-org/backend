package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/converter"
	"github.com/geekcamp-vol11-team30/backend/repository/internal/models"
	"github.com/geekcamp-vol11-team30/backend/util"
	"github.com/oklog/ulid/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type OauthRepository interface {
	RegisterProvider(ctx context.Context, op entity.OauthProvider) (entity.OauthProvider, error)
	FetchProviderByName(ctx context.Context, name string) (entity.OauthProvider, error)
	RegisterOauthUserInfo(ctx context.Context, oui entity.OauthUserInfo) (entity.OauthUserInfo, error)
	UpdateOauthUserInfo(ctx context.Context, oui entity.OauthUserInfo) (entity.OauthUserInfo, error)
	FetchOauthUserInfos(ctx context.Context, user entity.User) ([]entity.OauthUserInfo, error)
	FetchOauthUserInfo(ctx context.Context, providerId ulid.ULID, user entity.User) (entity.OauthUserInfo, error)
	FetchUserInfoByUid(ctx context.Context, providerId ulid.ULID, uid string) (*entity.OauthUserInfo, error)
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
	op.ID = util.GenerateULIDNow()
	opm := converter.OauthProviderEntityToModel(ctx, op)

	err := opm.Upsert(ctx, oar.db, boil.Infer(), boil.Infer())
	if err != nil {
		return entity.OauthProvider{}, err
	}
	newOp, err := converter.OauthProviderModelToEntity(ctx, opm)
	if err != nil {
		return entity.OauthProvider{}, fmt.Errorf("failed to convert OauthProviderModel to entity on RegisterProvider: %w", err)
	}
	return *newOp, nil
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
	oui.ID = util.GenerateULID(ctx)

	ouim := converter.OauthUserInfoEntityToModel(ctx, oui)

	err := ouim.Upsert(ctx, oar.db, boil.Infer(), boil.Infer())
	if err != nil {
		return entity.OauthUserInfo{}, err
	}

	newOui, err := converter.OauthUserInfoModelToEntity(ctx, ouim)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to convert OauthUserInfoModel to entity on RegisterOauthUserInfo: %w", err)
	}
	return newOui, nil
}

// UpdateOauthUserInfo implements OauthRepository.
func (oar *oauthRepository) UpdateOauthUserInfo(ctx context.Context, oui entity.OauthUserInfo) (entity.OauthUserInfo, error) {
	ouim := converter.OauthUserInfoEntityToModel(ctx, oui)

	_, err := ouim.Update(ctx, oar.db, boil.Infer())
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to update OauthUserInfoModel: %w", err)
	}

	newOui, err := converter.OauthUserInfoModelToEntity(ctx, ouim)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to convert OauthUserInfoModel to entity on UpdateOauthUserInfo: %w", err)
	}
	return newOui, nil
}

// FetchOauthUserInfo implements OauthRepository.
func (oar *oauthRepository) FetchOauthUserInfos(ctx context.Context, user entity.User) ([]entity.OauthUserInfo, error) {
	ouism, err := models.OauthUserInfos(
		models.OauthUserInfoWhere.UserID.EQ(util.ULIDToString(user.ID)),
		qm.Load(
			models.OauthUserInfoRels.Provider,
		),
	).All(ctx, oar.db)
	if err != nil {
		return nil, err
	}
	ouis, err := converter.OauthUserInfoSliceModelToEntity(ctx, ouism)
	if err != nil {
		return nil, fmt.Errorf("failed to convert OauthUserInfoModel to entity on FetchOauthUserInfos: %w", err)
	}
	return ouis, nil
	// ouis := make([]entity.OauthUserInfo, len(ouism))
	// for i, oui := range ouism {
	// 	ouii, err := converter.OauthUserInfoModelToEntity(ctx, oui)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("failed to convert OauthUserInfoModel to entity on FetchOauthUserInfos: %w", err)
	// 	}
	// 	ouis[i] = ouii
	// 	// id, _ := util.ULIDFromString(oui.ID)
	// 	// userId, _ := util.ULIDFromString(oui.UserID)
	// 	// providerId, _ := util.ULIDFromString(oui.ProviderID)
	// 	// ouis[i] = entity.OauthUserInfo{
	// 	// 	ID:                    id,
	// 	// 	UserId:                userId,
	// 	// 	ProviderId:            providerId,
	// 	// 	ProviderUid:           oui.ProviderUID,
	// 	// 	AccessToken:           oui.AccessToken,
	// 	// 	RefreshToken:          oui.RefreshToken,
	// 	// 	AccessTokenExpiresAt:  oui.AccessTokenExpiresAt,
	// 	// 	RefreshTokenExpiresAt: oui.RefreshTokenExpiresAt.Ptr(),
	// 	// }
	// }
	// return ouis, nil
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

	oui, err := converter.OauthUserInfoModelToEntity(ctx, ouim)
	if err != nil {
		return entity.OauthUserInfo{}, fmt.Errorf("failed to convert OauthUserInfoModel to entity: %w", err)
	}
	return oui, nil
}

func (oar *oauthRepository) FetchUserInfoByUid(ctx context.Context, providerId ulid.ULID, uid string) (*entity.OauthUserInfo, error) {
	ouim, err := models.OauthUserInfos(
		models.OauthUserInfoWhere.ProviderUID.EQ(uid),
		models.OauthUserInfoWhere.ProviderID.EQ(util.ULIDToString(providerId)),
	).One(ctx, oar.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	oui, err := converter.OauthUserInfoModelToEntity(ctx, ouim)
	if err != nil {
		return nil, fmt.Errorf("failed to convert OauthUserInfoModel to entity on FetchUserInfoByUid: %w", err)
	}
	return &oui, nil
}
