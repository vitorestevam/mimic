package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func GetRandomDate() string {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0).String()
}

func GetRandomUUID() string {
	b := make([]byte, 16)
	crypto.Read(b)

	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}
