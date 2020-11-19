package models

import "time"

// URL is model for hold original url and shorted url
type URL struct {
	Shorted   string
	Original  string
	CreatedAt time.Time
	ExpiredAt time.Time
}
