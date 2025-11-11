package main

import (
	"fmt"
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
			fmt.Println("Usage: btcpeek address <address>")
			return
		}
		c.LookupAddress(args[2])
	case "tx":
		if len(args) < 3 {
			fmt.Println("Usage: btcpeek tx <txid>")
			return
		}
		c.LookupTransaction(args[2])
	case "block":
		if len(args) < 3 {
			fmt.Println("Usage: btcpeek block <blockhash>")
			return
		}
		c.LookupBlock(args[2])
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage: btcpeek <command> [arguments]")
	fmt.Println("Commands:")
	fmt.Println("  address <address>   Lookup Bitcoin address information")
	fmt.Println("  tx <txid>          Lookup Bitcoin transaction information")
	fmt.Println("  block <blockhash>  Lookup Bitcoin block information")
}

func main() {
	run(os.Args)
}
