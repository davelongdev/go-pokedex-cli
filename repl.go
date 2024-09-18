package main

import (
	"bufio"
	"fmt"
	"os"
)

func startRepl() {
  scanner := bufio.NewScanner(os.Stdin)
  for {
    fmt.Printf("Pokedex > ")
    scanner.Scan()
    text := scanner.Text()
    fmt.Printf("echoing: %v\n", text)
  }
}
