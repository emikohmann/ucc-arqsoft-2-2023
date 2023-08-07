package domain

import "time"

type Flight struct {
	FlightNumber string    `json:"flight_number"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Departure    time.Time `json:"departure"`
	Arrival      time.Time `json:"arrival"`
	Availability int       `json:"availability"`
	Price        float64   `json:"price"`
}

type SearchResult struct {
	Total     int      `json:"total"`
	Results   []Flight `json:"results"`
	RequestID string   `json:"request_id"`
}
