package shorten

import (
	"math/rand"

	"github.com/mattkelly/url-shortener-go/db"
)

const slugLen uint = 5

func randomString(n uint) string {
	const charset string = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"

	// Strings are indexed as bytes in go: https://blog.golang.org/strings
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func Shorten(longUrl string) string {
	var shortUrl string

	exists := true
	for exists {
		shortUrl = randomString(slugLen)
		exists = db.Exists(shortUrl)
	}

	err := db.Set(shortUrl, longUrl)
	if err != nil {
		panic(err)
	}

	return shortUrl
}
