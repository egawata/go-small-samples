package sample

import (
	"crypto/rand"
	"log"
	"math/big"
)

func RandInt64() int64 {
	max := big.NewInt(100)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Fatal(err)
	}

	return n.Int64()
}
