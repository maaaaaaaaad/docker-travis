package main

import (
	"fmt"
	"math"
)

// d/dx sin(x) = lim(h->0) [sin(x+h) - sin(x)] / h
//            = lim(h->0) [sin(x)cos(h) + cos(x)sin(h) - sin(x)] / h
//            = lim(h->0) [cos(h) - 1] * sin(x) / h + cos(x) * lim(h->0) sin(h) / h
//            = cos(x)

// d/dx cos(x) = lim(h->0) [cos(x+h) - cos(x)] / h
//            = lim(h->0) [cos(x)cos(h) - sin(x)sin(h) - cos(x)] / h
//            = lim(h->0) [-sin(h)] * sin(x) / h + cos(x) * lim(h->0) [cos(h) - 1] / h
//            = -sin(x)

// d/dx tan(x) = d/dx [sin(x) / cos(x)]
//            = [cos(x) * d/dx sin(x) - sin(x) * d/dx cos(x)] / cos^2(x)
//            = [cos(x) * cos(x) - (-sin(x)) * sin(x)] / cos^2(x)
//            = sec^2(x)

// d/dx log_a(x) = 1 / (x * ln(a))
// d/dx ln(x) = 1 / x

func main() {
	sinX := func(x float64) float64 {
		return math.Sin(x)
	}
	cosX := differentiateTrigFunction(sinX, 0.5)
	fmt.Printf("sin(0.5)의 도함수는 %.4f입니다.\n", cosX)

	tanX := func(x float64) float64 {
		return math.Tan(x)
	}
	sec2X := differentiateTrigFunction(tanX, 0.5)
	fmt.Printf("tan(0.5)의 도함수는 %.4f입니다.\n", sec2X)

	logX := func(x float64) float64 {
		return math.Log2(x)
	}
	dLogX := differentiateLogFunction(logX, 10, 2)
	fmt.Printf("log_2(10)의 도함수는 %.4f입니다.\n", dLogX)

	lnX := func(x float64) float64 {
		return math.Log(x)
	}
	dLnX := differentiateLnFunction(lnX, 3)
	fmt.Printf("ln(3)의 도함수는 %.4f입니다.\n", dLnX)
}

func differentiateTrigFunction(f func(x float64) float64, x float64) float64 {
	h := 0.00000001
	return (f(x+h) - f(x)) / h
}

func differentiateLogFunction(f func(x float64) float64, x float64, a float64) float64 {
	h := 0.00000001
	return 1 / (x*math.Log(a) + h)
}

func differentiateLnFunction(f func(x float64) float64, x float64) float64 {
	h := 0.00000001
	return 1 / (x + h)
}
