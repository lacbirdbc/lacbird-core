package main

import (
	"fmt"
	"log"

	"github.com/lacbirdbc/lacbird-core/account"
)

func main() {
	a, err := account.NewAccount()
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Printf("%#v\n", a)
	fmt.Println("prvKey", a.PrivateKey.D)
}
