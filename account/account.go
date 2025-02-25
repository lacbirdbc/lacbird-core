package account

import (
	"crypto/ecdsa"
	"crypto/rand"
	"os"

	"github.com/dustinxie/ecc"
	"github.com/lacbirdbc/lacbird-core/address"
)

type Account struct {
	Address    address.Address
	PrivateKey *ecdsa.PrivateKey
}

func NewAccount() (Account, error) {
	prv, err := ecdsa.GenerateKey(ecc.P256k1(), rand.Reader)
	if err != nil {
		return Account{}, err
	}
	addr := address.NewAddress(&prv.PublicKey)
	return Account{Address: addr, PrivateKey: prv}, nil
}

func ReadAccount(path string, pass []byte) (Account, error) {
	cprv, err := os.ReadFile(path)
	if err != nil {
		return Account{}, nil
	}

	jprv, err := decryptWithPassword(cprv, pass)
	if err != nil {
		return Account{}, err
	}
	return Account{
		Address:    address.Address(jprv),
		PrivateKey: nil,
	}, nil
}

func decryptWithPassword(ciph, pass []byte) ([]byte, error) {
	return nil, nil
}
