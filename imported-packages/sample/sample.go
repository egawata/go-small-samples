package sample

import (
	"bytes"
	"crypto/rand"
	"log"
	"math/big"
)

func RandInt64() int64 {
	rd := bytes.NewBufferString("hoge")
	max := big.NewInt(100)
	n, err := rand.Int(rd, max)
	if err != nil {
		log.Fatal(err)
	}

	return n.Int64()
}
