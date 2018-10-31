package liteurl

import (
	"crypto/md5"
	"fmt"
	"math"
	"strings"
)

const (
	baseStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+-"
)

var base uint64

func init() {
	base = uint64(len(baseStr))
}

func GetMD5(src string) string {
	hasher := md5.New()
	hasher.Write([]byte(src))
	r := hasher.Sum(nil)
	return fmt.Sprintf("%x", r)
}

func ConvertToBase(input uint64) string {
	var r []byte
	for input > 0 {
		r = append(r, baseStr[input%base])
		input /= base
	}
	return string(r)
}

func RestoreBase10(src string) uint64 {
	var a uint64
	array := []byte(src)
	for i := 0; i < len(array); i++ {
		var power uint64 = uint64(math.Pow(float64(base), float64(i)))
		base := strings.Index(baseStr, string(array[i]))
		a += uint64(base) * power
	}
	return a
}
