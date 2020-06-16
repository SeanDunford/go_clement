package main

import (
	"math/rand"
)

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomName(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += RandomString(rand.Intn(9) + 1)
	}

	return s
}
