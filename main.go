package main

import (
	"os"

	"github.com/daut/btcpeek/clients"
	"github.com/daut/btcpeek/commands"
)

func run(args []string) {
	cmd := args[1]

	c := commands.NewCommands(clients.FetchData)

	switch cmd {
	case "address":
		if len(args) < 3 {
			println("Usage: btcpeek address <address>")
			return
		}
		c.LookupAddress(args[2])
	case "tx":
		if len(args) < 3 {
			println("Usage: btcpeek tx <txid>")
			return
		}
		c.LookupTransaction(args[2])
	case "block":
		if len(args) < 3 {
			println("Usage: btcpeek block <blockhash>")
			return
		}
		c.LookupBlock(args[2])
	}
}

func main() {
	run(os.Args)
}
