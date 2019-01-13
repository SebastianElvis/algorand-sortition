package sortition

import (
	"github.com/coniks-sys/coniks-go/crypto/vrf"
)

func Sortition(sk vrf.PrivateKey, seed []byte, threshold int64, role []byte, balance int64, total int64) ([]byte, []byte, int64) {
	hash, proof := sk.Prove(append(seed, role...))
	p := float64(threshold) / float64(total)
	j := int64(0)
	normalizedHash := normalizeHash(hash)

	// main loop for computing j
	for {
		lower := accB(balance, p, j)
		higher := accB(balance, p, j+1)
		if lower <= normalizedHash && normalizedHash < higher {
			break
		}
		j++
	}
	return hash, proof, j
}

func Verify(pk vrf.PublicKey, hash []byte, proof []byte, seed []byte, threshold int64, role []byte, balance int64, total int64, weight int64) bool {
	if !pk.Verify(append(seed, role...), hash, proof) {
		return false
	}
	p := float64(threshold) / float64(total)
	j := int64(0)
	normalizedHash := normalizeHash(hash)
	// main loop for computing j
	for {
		lower := accB(balance, p, j)
		higher := accB(balance, p, j+1)
		if lower <= normalizedHash && normalizedHash < higher {
			break
		}
		j++
	}
	if j != weight {
		return false
	}
	return true
}
