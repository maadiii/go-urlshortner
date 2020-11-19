package interactors

import (
	"strings"
	"time"
	"urlshortner/domain/interactors"
	"urlshortner/domain/models"
	"urlshortner/infrastructure/utils"
	ui "urlshortner/usecase/interactors"
	"urlshortner/usecase/repositories"
)

type url struct {
	repository repositories.URL
}

func NewUrl(repository repositories.URL) ui.URL {
	return url{repository}
}

func (self url) Add(url *interactors.OriginalUrl) (interactors.ShortedUrl, error) {
	var shorted interactors.ShortedUrl
	// TODO: get length from config file
	model := models.URL{
		Shorted:   strings.ToLower(utils.RandomBase64String(7)),
		Original:  url.Url,
		CreatedAt: time.Now(),
		ExpiredAt: time.Now().Add(2160 * time.Hour),
	}

	err := self.repository.Create(&model)
	if err != nil {
		return shorted, err
	}

	shorted = interactors.ShortedUrl{
		Url:       model.Shorted,
		CreatedAt: model.CreatedAt,
		ExpiredAt: model.ExpiredAt,
	}

	return shorted, nil
}

func (self url) Get(shortedUrl interactors.ShortedUrl) (interactors.CompleteUrl, error) {
	var completeUrl interactors.CompleteUrl
	model := models.URL{Shorted: shortedUrl.Url}
	err := self.repository.Read(&model)
	if err != nil {
		return completeUrl, err
	}

	completeUrl = interactors.CompleteUrl{
		OriginalUrl: model.Original,
		ShortedUrl:  model.Shorted,
		CreatedAt:   model.CreatedAt,
		ExpirtedAt:  model.ExpiredAt,
	}

	return completeUrl, nil
}
