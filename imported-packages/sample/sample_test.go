package sample_test

import (
	"testing"

	"github.com/egawata/go-small-samples/imported-packages/sample"
)

func TestRandInt64(t *testing.T) {
	v := sample.RandInt64()
	if v > 100 {
		t.Fatal("too much")
	}
}
