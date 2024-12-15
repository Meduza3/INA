package main

import (
	"dijkstra/dimacs/lexer"
	"dijkstra/dimacs/parser"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	baseLexer := lexer.NewScanner()
	baseLexer.In = file
	lexer_wrapper := lexer.NewLexerWrapper(baseLexer)
	parser := parser.YyNewParser()

	parser.Parse(lexer_wrapper)
	fmt.Println(parser.GetGraph())
}
