package main

import (
	"os"
	"strconv"

	"github.com/ProggerX/pabc-flowchart/internal/extra"
	"github.com/ProggerX/pabc-flowchart/internal/parser"
	"github.com/charmbracelet/log"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 && len(args) != 2 {
		log.Fatal("Expected only one argument - filename, or two arguments - filename and verbose level")
	}
	filename := args[0]
	if len(args) == 2 {
		vl, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal("Verbose level must be a number")
		}
		switch vl {
		case 1:
			log.SetLevel(log.FatalLevel)
		case 2:
			log.SetLevel(log.ErrorLevel)
		case 3:
			log.SetLevel(log.WarnLevel)
		case 4:
			log.SetLevel(log.InfoLevel)
		case 5:
			log.SetLevel(log.DebugLevel)
		}
	}
	lines, err := extra.ReadLines(filename)
	if err != nil {
		log.Fatal("Error while reading file:", "err", err)
	}
	out := parser.ParseFile(lines)
	err = extra.WriteLines("a.out.md", out)
	if err != nil {
		log.Fatal("Error while writing result:", "err", err)
	}
	log.Info("Success! Mermaid flowchart can be found in a.out.md file")
}
