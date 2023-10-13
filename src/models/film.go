package models

import "time"

type Film struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rate        float64   `json:"rate"`
	Genre       string    `json:"genre"`
	ReleaseDT   time.Time `json:"release_dt"`
	Duration    uint      `json:"duration"`
	CountryID   uint64    `json:"country_id"`
}
