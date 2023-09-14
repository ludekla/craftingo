package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func run(msg string) {
	fmt.Println(msg)
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
	} else {
		runPrompt()
	}

}
