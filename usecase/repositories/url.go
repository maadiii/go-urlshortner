package repositories

import "urlshortner/domain/models"

type URL interface {
	Create(url *models.URL) error
	Read(url *models.URL) error
}
