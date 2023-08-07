package services

import (
	"context"
	"flights-api/domain"
	"time"
)

const (
	dateLayout = "2006-01-02"
)

func SearchFlights(ctx context.Context) domain.SearchResult {
	departure, _ := time.Parse(dateLayout, "2024-04-01")
	arrival, _ := time.Parse(dateLayout, "2024-05-01")

	return domain.SearchResult{
		Total: 3,
		Results: []domain.Flight{
			{
				FlightNumber: "F001",
				Origin:       "EZE",
				Destination:  "MAD",
				Departure:    departure,
				Arrival:      arrival,
				Availability: 10,
				Price:        500000,
			},
			{
				FlightNumber: "F002",
				Origin:       "EZE",
				Destination:  "MAD",
				Departure:    departure,
				Arrival:      arrival.Add(2 * time.Hour),
				Availability: 5,
				Price:        450000,
			},
			{
				FlightNumber: "F003",
				Origin:       "EZE",
				Destination:  "MAD",
				Departure:    departure,
				Arrival:      arrival.Add(4 * time.Hour),
				Availability: 3,
				Price:        300000,
			},
		},
	}
}
