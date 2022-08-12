package main

import (
	"fmt"
)

func help(args []string) {
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("  sarasvati <command> [<args>]")
	fmt.Println("")
	fmt.Println("Available commands:")
	fmt.Println("")
	fmt.Println("  work")
	fmt.Println("  auto")
	fmt.Println("  index")
	fmt.Println("    show <ID>")
	fmt.Println("    map <ID>")
	fmt.Println("  serve")
	fmt.Println("  help")
	fmt.Println("")
	fmt.Println("Use `sarasvati help <command>` for more information about a command.")
	fmt.Println("")

}
