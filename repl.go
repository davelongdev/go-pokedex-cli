package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func cleanInput(str string) []string {
  lowered := strings.ToLower(str)
  words := strings.Fields(lowered)
  return words
}
