package utils

import (
	"encoding/base64"
	"math"
	"math/rand"
)

func RandomBase64String(length int) string {
	buff := make([]byte, int(math.Round(float64(length)/float64(1.333333333333))))
	rand.Read(buff)
	str := base64.RawURLEncoding.EncodeToString(buff)

	return str[:length]
}
