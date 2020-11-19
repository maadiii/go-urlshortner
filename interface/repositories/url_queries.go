package repositories

const (
	INSERT_URL = `
			INSERT INTO urls (shorted, original, created_at, expired_at)
    	VALUES ($1, $2, $3, $4)
		`
	READ_URL = `
			SELECT original, created_at, expired_at FROM urls
			WHERE shorted = $1 AND expired_at > $2
    `
)
