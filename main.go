package main

import (
	"fmt"
	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
)

func main() {
	chain := &chaincfg.TestNet3Params

	addrStr := "qz32sqj8jwjfs8uxa36anlnx747fln0drs4n42hu3p"
	addr, err := bchutil.DecodeAddress(addrStr, chain)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(addr.String())

	oldaddr, err := bchutil.NewLegacyAddressPubKeyHash(addr.ScriptAddress(), chain)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	fmt.Println(oldaddr.String())

	newaddr, err := bchutil.NewAddressPubKeyHash(oldaddr.ScriptAddress(), chain)
	fmt.Println(newaddr.String())
}
