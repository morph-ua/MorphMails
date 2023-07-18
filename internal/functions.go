package main

import (
	"math/rand"
)

//goland:noinspection ALL
func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func syncWithClients(json FinalResult, url string) int {
	r, err := req.R().
		SetHeader("Content-type", "application/json").
		SetBodyJsonMarshal(json).
		SetHeader("user-agent", "github.com/voxelin").
		Post(url)

	if err != nil || r.StatusCode != 200 {
		return 403
	}

	return 0
}
