package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

const LengthUniqueString = 10

func UniqueString() string {
	rand.Seed(time.Now().UnixNano())
	newByte := make([]byte, LengthUniqueString)
	rand.Read(newByte)
	return fmt.Sprintf("%x", newByte)[:LengthUniqueString]
}
