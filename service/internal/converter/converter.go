package converter

import (
	"fmt"
	"time"

	"github.com/geekcamp-vol11-team30/backend/entity"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
)

func OauthUserInfoEntityToOauth2Token(oui entity.OauthUserInfo) *oauth2.Token {
	return &oauth2.Token{
		AccessToken:  oui.AccessToken,
		RefreshToken: oui.RefreshToken,
		Expiry:       oui.AccessTokenExpiresAt,
	}
}

func CalendarEventsToEntity(events *calendar.Events) ([]entity.CalendarEvent, error) {
	ce := make([]entity.CalendarEvent, len(events.Items))
	for i, e := range events.Items {
		event, err := CalendarEventToEntity(e)
		if err != nil {
			return nil, fmt.Errorf("failed to convert CalendarEvent to entity on CalendarEventsToEntity: %w", err)
		}
		ce[i] = event
	}
	return ce, nil
}
func CalendarEventToEntity(event *calendar.Event) (entity.CalendarEvent, error) {
	start, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}
	end, err := time.Parse(time.RFC3339, event.End.DateTime)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}

	return entity.CalendarEvent{
		Name:      event.Summary,
		StartTime: start,
		EndTime:   end,
	}, nil
}
