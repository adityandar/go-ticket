package structs

import "time"

type Event struct {
	Id          string    `json:"id"`
	OrganizerId string    `json:"organizer_id"`
	Title       string    `json:"title"`
	DateTime    time.Time `json:"date_time"`
	CreatedAt   time.Time `json:"created_at"`
}
