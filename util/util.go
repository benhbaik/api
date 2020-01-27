package util

import (
	"math/rand"
	"path"
	"strings"
	"time"
)

// ShiftPath Can remove first two chunks of URL path.
// Input: "/api/user/1234/"
// Output: "/user/1234", "/1234"
func ShiftPath(p string) (head, tail string) {
	p = path.Clean("/" + p)
	i := strings.Index(p[1:], "/") + 1
	// if no "/" left in path return "/"
	if i <= 0 {
		return p[1:], "/"
	}
	return p[1:i], p[i:]
}

// RandStringRunes Generates a random string of an argument specified length.
func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
