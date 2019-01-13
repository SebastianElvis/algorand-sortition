package sortition

import (
	"testing"

	"github.com/coniks-sys/coniks-go/crypto/vrf"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	sk, err := vrf.GenerateKey(nil)
	assert.NoError(t, err)
	pk, succ := sk.Public()
	assert.Equal(t, true, succ)

	seed := []byte{1, 2, 3, 4, 5}
	role := []byte{7, 7, 7, 7, 7, 7, 7}
	threshold := int64(40)
	balance := int64(3)
	total := int64(100)

	hash, proof, j := Sortition(sk, seed, threshold, role, balance, total)
	t.Log("Value of j is", j)
	valid := Verify(pk, hash, proof, seed, threshold, role, balance, total, j)
	assert.Equal(t, true, valid)
}
