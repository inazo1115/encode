package encode

import (
	"math"
	"strings"
)

func ToUnary(x int) string {
	if x < 1 {
		panic("input must be natural number")
	}
	return strings.Repeat("0", x-1) + "1"
}

func ToBinary(x int) string {
	if x < 0 {
		panic("input must be nonnegative integer")
	}
	ret := make([]string, lenBinary(x))
	for i := len(ret) - 1; i >= 0; i-- {
		if x%2 == 0 {
			ret[i] = "0"
		} else {
			ret[i] = "1"
		}
		x = x / 2
	}
	return strings.Join(ret, "")
}

func lenBinary(x int) int {
	m := float64(x)
	return int(math.Floor(math.Max(math.Log2(m)+1, 1)))
}

func ToGamma(x int) string {
	if x < 1 {
		panic("input must be natural number")
	}
	binary := ToBinary(x)
	return ToUnary(len(binary)) + binary[1:]
}

func ToDelta(x int) string {
	if x < 1 {
		panic("input must be natural number")
	}
	binary := ToBinary(x)
	return ToGamma(len(binary)) + binary[1:]
}

func ToGolomb(x int, k int) string {
	if x < 1 {
		panic("input must be natural number")
	}
	if k < 0 {
		panic("k must be nonnegative integer")
	}
	x_ := x - 1
	r := (x_ / k) + 1
	m := x_ % k
	l := len(ToBinary(k - 1))
	b := ToBinary(m)
	return ToUnary(r) + strings.Repeat("0", l - len(b)) + b
}
