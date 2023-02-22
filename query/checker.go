package query

import (
	"github.com/sandertv/go-raknet"
)

func CheckServer(address string) bool {
	_, err := raknet.Dial(address)

	if err != nil {
		return false
	} else {
		return true
	}
}