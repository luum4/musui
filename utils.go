package main

import (
	"math/rand"
	"strconv"
)

func RandomInt(min, max int) string {
    c := min + rand.Intn(max-min)
	return strconv.Itoa(c)
}