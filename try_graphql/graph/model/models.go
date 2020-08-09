package model

import "time"

type Video struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	UserID      int           `json:"-"`
	URL         string        `json:"url"`
	CreatedAt   time.Time     `json:"createdAt"`
	Screenshots []*Screenshot `json:"screenshots"`
	Related     []*Video      `json:"related"`
}
