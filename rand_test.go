package main

import (
	"fmt"
	"testing"
)

func TestGaussianRand(t *testing.T) {
	gsRand := GaussianRand(60.0, 4.5)
	for i := 0; i < 10; i++ {
		f := gsRand()
		fmt.Println(f)
	}
}
