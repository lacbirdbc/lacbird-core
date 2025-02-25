package secp256k1

import (
	"crypto/ecdsa"
	"math/big"
)

type P256k1PublicKey struct {
	Curve string   `json:"curve"`
	X     *big.Int `json:"x"`
	Y     *big.Int `json:"y"`
}

func NewP256k1PublicKey(pub *ecdsa.PublicKey) P256k1PublicKey {
	return P256k1PublicKey{Curve: "P-256k1", X: pub.X, Y: pub.Y}
}
