package server

import "math"

func isPrime(p int64) bool {
	max := 1 + int64(math.Sqrt(float64(p)))
	for i := int64(3); i < max; i += 2 {
		if p%i == 0 {
			return false
		}
	}
	return true
}
