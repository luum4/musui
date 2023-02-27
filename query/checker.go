package query

import (
	"github.com/sandertv/go-raknet"
)

func CheckServer(address string) bool {
	_, err := raknet.Ping(address)

	if err != nil {
		return false
	}

	return true
}
