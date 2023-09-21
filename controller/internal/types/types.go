package types

import (
	"time"

	"github.com/geekcamp-vol11-team30/backend/entity"
)

type ExternalEventRequest struct {
	TimeMin *time.Time `query:"timeMin"`
	TimeMax *time.Time `query:"timeMax"`
}
type UserResponse struct {
	entity.User
	Providers []ProviderResponse `json:"providers"`
}
type ProviderResponse struct {
	Name       string `json:"name"`
	Registered bool   `json:"registered"`
}

func NewProviderResponse(ops []entity.OauthProvider, ouis []entity.OauthUserInfo) []ProviderResponse {
	var res []ProviderResponse
	for _, op := range ops {
		var registered bool
		for _, oui := range ouis {
			if op.ID == oui.ProviderId {
				registered = true
			}
		}
		res = append(res, ProviderResponse{
			Name:       op.Name,
			Registered: registered,
		})
	}
	return res
}
