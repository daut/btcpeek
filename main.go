package main

import (
	"os"

	"github.com/daut/btcpeek/clients"
	"github.com/daut/btcpeek/commands"
)

func main() {
	cmd := os.Args[1]

	c := commands.NewCommands(clients.FetchData)

	switch cmd {
	case "address":
		if len(os.Args) < 3 {
			println("Usage: btcpeek address <address>")
			return
		}
		c.LookupAddress(os.Args[2])
	case "tx":
		if len(os.Args) < 3 {
			println("Usage: btcpeek tx <txid>")
			return
		}
		c.LookupTransaction(os.Args[2])
	case "block":
		if len(os.Args) < 3 {
			println("Usage: btcpeek block <blockhash>")
			return
		}
		c.LookupBlock(os.Args[2])
	}
}
