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

func GoogleCalendarEventsToEntity(events []*calendar.Event, calendarID string, calendarName string) (entity.Calendar, error) {
	ce := make([]entity.CalendarEvent, len(events))
	for i, e := range events {
		event, err := GoogleCalendarEventToEntity(e)
		if err != nil {
			return entity.Calendar{}, fmt.Errorf("failed to convert CalendarEvent to entity on CalendarEventsToEntity: %w", err)
		}
		ce[i] = event
	}

	return entity.Calendar{
		Events:       ce,
		Provider:     "google",
		CalendarName: calendarName,
		CalendarID:   calendarID,
		Count:        len(events),
	}, nil
}
func GoogleCalendarEventToEntity(event *calendar.Event) (entity.CalendarEvent, error) {
	if event.Start.DateTime == "" {
		return parseGoogleAllDayEventToEntity(event)
	}
	return parseGoogleNormalEventToEntity(event)

}

func parseGoogleAllDayEventToEntity(event *calendar.Event) (entity.CalendarEvent, error) {
	url := event.HtmlLink
	displayOnly := event.Transparency == "transparent"
	startDate, err := time.Parse("2006-01-02", event.Start.Date)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}
	endDate, err := time.Parse("2006-01-02", event.End.Date)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}
	return entity.CalendarEvent{
		Name:        event.Summary,
		StartDate:   &entity.Date{Time: startDate},
		EndDate:     &entity.Date{Time: endDate},
		IsAllDay:    true,
		URL:         url,
		DisplayOnly: displayOnly,
	}, nil
}

func parseGoogleNormalEventToEntity(event *calendar.Event) (entity.CalendarEvent, error) {
	url := event.HtmlLink
	displayOnly := event.Transparency == "transparent"
	fmt.Printf("event.Transparency: %+v\n", event.Transparency)
	start, err := time.Parse(time.RFC3339, event.Start.DateTime)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}
	end, err := time.Parse(time.RFC3339, event.End.DateTime)
	if err != nil {
		return entity.CalendarEvent{}, fmt.Errorf("failed to parse time: %w", err)
	}
	return entity.CalendarEvent{
		Name:        event.Summary,
		StartTime:   &start,
		EndTime:     &end,
		IsAllDay:    false,
		URL:         url,
		DisplayOnly: displayOnly,
	}, nil
}
