package address

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"

	"github.com/lacbirdbc/lacbird-core/crypto/secp256k1"
	"golang.org/x/crypto/sha3"
)

type Address string

func NewAddress(pub *ecdsa.PublicKey) Address {
	jpub, _ := json.Marshal(secp256k1.NewP256k1PublicKey(pub))
	hash := make([]byte, 64)
	sha3.ShakeSum256(hash, jpub)
	return Address(hex.EncodeToString(hash[:32]))
}
