package interactors

import "urlshortner/domain/interactors"

type URL interface {
	Add(originalUrl *interactors.OriginalUrl) (interactors.ShortedUrl, error)
	Get(shortedUrl interactors.ShortedUrl) (interactors.CompleteUrl, error)
}
