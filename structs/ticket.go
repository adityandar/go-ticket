package structs

import "time"

type Ticket struct {
	Id         string    `json:"id"`
	EventId    string    `json:"event_id"`
	AudienceId string    `json:"audience_id"`
	FullName   string    `json:"full_name"`
	CreatedAt  time.Time `json:"created_at"`
}
