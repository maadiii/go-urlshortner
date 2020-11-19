package datastore

import (
	"database/sql"
	"fmt"
	"urlshortner/infrastructure/config"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

// Session sql and nosql databases session
type Session struct {
	Postgres postgresSession
}

// NewSession creates and returns session
func NewSession() (*Session, error) {
	sqlSession, err := newSQLSession()
	if err != nil {
		return nil, err
	}

	return &Session{*sqlSession}, nil
}

type postgresSession struct {
	*sql.DB
}

func newSQLSession() (*postgresSession, error) {
	conf, err := getDBConfig()
	if err != nil {
		return nil, err
	}

	session, err := getSQLSession(conf.Databases.Postgres.Driver, conf.Databases.Postgres.URL)
	if err != nil {
		return nil, err
	}

	return &postgresSession{session}, nil
}

func getSQLSession(driverName, url string) (*sql.DB, error) {
	session, err := sql.Open(driverName, url)
	if err != nil {
		return nil, errors.New(
			fmt.Sprintf(
				"%s %s %s %v", "session unable connect to", driverName, "database", err,
			),
		)
	}

	return session, nil
}

func getDBConfig() (*config.DataStoreConfig, error) {
	c, err := config.ConfigFactory(config.DATASTORE_CONFIG)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	conf := c.(*config.DataStoreConfig)

	return conf, nil
}

// NewTestSession creates and returns session for test goals
func NewTestSession() (*Session, error) {
	sqlSession, err := newSQLTestSession()
	if err != nil {
		return nil, err
	}

	return &Session{*sqlSession}, nil
}

func newSQLTestSession() (*postgresSession, error) {
	conf, err := getDBConfig()
	if err != nil {
		return nil, err
	}

	session, err := getSQLSession(conf.Databases.Postgres.Driver, conf.Databases.Postgres.Test)
	if err != nil {
		return nil, err
	}

	return &postgresSession{session}, nil
}
