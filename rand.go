package lib

import (
	"math/rand"
	"time"
)

const An = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandAlphaNum(n int, f bool) string {
	r := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	for i := range r {
		r[i] = An[rand.Intn(len(An))]
	}
	return string(r)
}
