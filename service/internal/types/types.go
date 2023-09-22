package types

type MsUser struct {
	Id            string `json:"id"`
	DisplayName   string `json:"displayName"`
	PrincipalName string `json:"userPrincipalName"`
}

type MsCalendar struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	IsDefaultCalendar bool   `json:"isDefaultCalendar"`
}
type MsCalendarResponse struct {
	Value []MsCalendar `json:"value"`
}

type MsCalendarEvent struct {
	Subject string `json:"subject"`
	Start   struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"start"`
	End struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"end"`
}

type MsCalendarEventsResponse struct {
	Value []MsCalendarEvent `json:"value"`
}
