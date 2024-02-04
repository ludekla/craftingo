package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	tk "golox/pkg/tokens"
)

func run(msg string) {
	scanner := tk.NewScanner(msg)
	scanner.ScanTokens()

	for i, token := range scanner.Tokens() {
		fmt.Printf("%d: %+v\n", i, token)
	}
}

func runPrompt() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print("> ")
	for sc.Scan() {
		msg := sc.Text()
		run(msg)
		fmt.Print("> ")
	}
}

func runFile(filename string) error {
	fp, err := os.Open(filename)
	if err != nil {
		log.Fatalf("runFile: %v", err)
	}
	defer fp.Close()
	sc := bufio.NewScanner(fp)
	for sc.Scan() {
		msg := sc.Text()
		run(msg)
	}
	return nil
}

func main() {

	fmt.Println("Hello World!")

	if len(os.Args) > 2 {
		fmt.Println("Usage: bin/main [OPTIONAL: filename]")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
		token := "uber"
		fmt.Printf("%s\n", token[0:1])
	} else {
		runPrompt()
	}

}
