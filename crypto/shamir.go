package crypto

import (
	"math/big"
	"math/rand"
)

func evaluatePolynomial(coefficients []*big.Int, x int) *big.Int {
	return big.NewInt(0)
}

// Generates n key shares (xi, P(xi)) where xi is [1...n] and P(xi) is a 256-bit integer
// t - 1 represents the degree of the polynomial P used to generate the n key shares
// which makes only t of them necessary to calculate K key used to encrypt the content
func GenKeyShares(secret [32]byte, t, n int) ([][32]byte, error) {
	// 1. Random selection of coeficcients
	coefficients := make([]*big.Int, t-1)
	for i := 0; i < t-1; i++ {
		randomCoefficient := make([]byte, 32)
		if _, err := rand.Read(randomCoefficient); err != nil {
			return nil, err
		}
		coefficient := big.NewInt(0)
		coefficient.SetBytes(randomCoefficient)
		coefficients[i] = coefficient
	}

	key := big.NewInt(0)
	key.SetBytes(secret[:])

	result := make([][32]byte, n)
	for i := 0; i < n; i++ {
		evaluation := evaluatePolynomial(coefficients, i)
		copy(result[i][:], evaluation.Bytes())
	}

	return result, nil
}