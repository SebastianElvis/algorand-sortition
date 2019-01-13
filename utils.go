package sortition

import (
	"math"
	"math/big"
)

func factorial(n int64) int64 {
	if n <= 1 {
		return int64(1)
	}
	return n * factorial(n-1)
}

// Compute the Binomial Distribution B(k; w, p)
// Assume you flip a coin for w times. Flip it with the possibility of the front face = p
// B(k; w, p) is the probability that front face occurs for exactly k times
func b(k, w int64, p float64) float64 {
	comb := float64(factorial(w) / factorial(k) / factorial(w-k))
	px := math.Pow(p, float64(k))
	qnx := math.Pow(1.0-p, float64(w-k))
	return comb * px * qnx
}

// B(0, w, p) + ... + B(j, w, p)
func accB(w int64, p float64, j int64) float64 {
	accB := float64(0)
	for k := int64(0); k <= j; k++ {
		accB += b(k, w, p)
	}
	return accB
}

// hash / 1^hashlen
func normalizeHash(hash []byte) float64 {
	maxInt256Array := make([]byte, 32)
	for i, _ := range maxInt256Array {
		maxInt256Array[i] = 255
	}
	return divide(hash, maxInt256Array)
}

func divide(a, b []byte) float64 {
	aInt, bInt := new(big.Int), new(big.Int)
	aInt.SetBytes(a[:])
	bInt.SetBytes(b[:])

	// convert to float
	aFloat, bFloat := new(big.Float).SetInt(aInt), new(big.Float).SetInt(bInt)
	result, _ := new(big.Float).Quo(aFloat, bFloat).Float64()
	return result
}
