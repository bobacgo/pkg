package main

import (
	"fmt"
	"strconv"
)

func Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func Decimal32(value float32) float32 {
	v, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 32)
	return float32(v)
}
