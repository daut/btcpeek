package main

import (
	"github.com/daut/btcpeek/commands"
	"os"
)

func main() {
	command := os.Args[1]

	switch command {
	case "address":
		if len(os.Args) < 3 {
			println("Usage: btcpeek address <address>")
			return
		}
		commands.LookupAddress(os.Args[2])
	case "tx":
		if len(os.Args) < 3 {
			println("Usage: btcpeek tx <txid>")
			return
		}
		commands.LookupTransaction(os.Args[2])
	}
}
