package main

import (
	"fmt"

	"github.com/digitalbitsorg/go/clients/frontier"
)

func main() {
	blob := "AAAAAH0t3PCq/ZCAvioV7ThK3edEcX3PcUrhoc0vngMsqGGWAAAAZAA1fpcAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAANg+Af8x49GLeB3Un+Qr9ESJJAnsZmqTNpOqM9dTDnhkAAAAAAAAAAAAPQkAAAAAAAAAAASyoYZYAAABA/L7Du9hlYzCtR9F89Mp9/alkCXsq9CWuJ1Mpql+Q16fHE4P2+H62p4cx+b2YUp/fUX73ucW+RPxOgSXmeV6uBQ=="

	resp, err := frontier.DefaultTestNetClient.SubmitTransaction(blob)
	if err != nil {
		panic(err)
	}

	fmt.Println("transaction posted in ledger:", resp.Ledger)
}
