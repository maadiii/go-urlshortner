package repositories

import (
	"database/sql"
	"time"
	"urlshortner/domain/models"
	"urlshortner/infrastructure/application"
	"urlshortner/infrastructure/datastore"
	"urlshortner/usecase/repositories"
)

type url struct {
	datastore.Session
}

func NewUrl(session datastore.Session) repositories.URL {
	return url{session}
}

func (self url) Create(url *models.URL) error {
	_, err := self.Postgres.Exec(INSERT_URL, url.Shorted, url.Original, url.CreatedAt, url.ExpiredAt)

	return err
}

func (self url) Read(url *models.URL) error {
	row := self.Postgres.QueryRow(READ_URL, url.Shorted, time.Now())
	err := row.Scan(&url.Original, &url.CreatedAt, &url.ExpiredAt)
	if err == sql.ErrNoRows {
		return application.NewErrNotFound(url.Shorted)
	}

	return err
}
