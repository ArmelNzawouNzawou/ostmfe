package event

import "time"

type Event struct {
	Id   string    `json:"id"`
	Name string    `json:"name"`
	Date time.Time `json:"date"`
}

type EventImage struct {
	ImageId     string `json:"imageId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventPartener struct {
	PartenerId  string `json:"partenerId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventPlace struct {
	PlaceId     string `json:"placeId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}

type EventProject struct {
	ProjectId   string `json:"projectId"`
	EventId     string `json:"eventId"`
	Description string `json:"description"`
}
