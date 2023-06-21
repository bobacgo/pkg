package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// GaussianRand
// 正态分布的参数
// mean 均值
// stddev 标准差
func GaussianRand(mean, stddev float64) func() float64 {
	// 初始化随机数生成器
	rand.NewSource(time.Now().UnixNano())
	return func() float64 {
		// 生成高斯随机数
		gaussian := rand.NormFloat64()*stddev + mean
		fStr := fmt.Sprintf("%.2f", gaussian)
		float, _ := strconv.ParseFloat(fStr, 64)
		return float
	}
}

// 生成随机序列（包括：大小写字母、数字）
func RandSeqID(n int) func() string {
	letters := []rune("0123456789abcdefghijklmnopgrstuvwxyzABCDEFGHIJKLMNOPGRSTUVWXYZ")
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() string {
		b := make([]rune, n)
		for i := range b {
			b[i] = letters[r.Intn(len(letters))]
		}
		return string(b)
	}
}

// RandNumber 指定范围的随机数
func RandNumber(start, end int) func() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return func() int {
		if start >= end {
			return end
		}
		num := r.Intn(end-start) + start
		return num
	}
}

func UUID() string {
	newId := strings.ReplaceAll(uuid.NewString(), "-", "")
	return newId
}
