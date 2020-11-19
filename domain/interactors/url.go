package interactors

import "time"

type OriginalUrl struct {
	Url string `json:"url"`
}

type ShortedUrl struct {
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	ExpiredAt time.Time `json:"expireAt"`
}

type CompleteUrl struct {
	OriginalUrl string    `json:"originalUrl"`
	ShortedUrl  string    `json:"shortedUrl"`
	CreatedAt   time.Time `json:"createdAt"`
	ExpirtedAt  time.Time `json:"expirtedAt"`
}
