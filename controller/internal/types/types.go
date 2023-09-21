package types

import "time"

type ExternalEventRequest struct {
	TimeMin *time.Time `query:"timeMin"`
	TimeMax *time.Time `query:"timeMax"`
}
