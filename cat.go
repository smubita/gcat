package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func unused(x ...interface{}) {}

func main() {
	var args []string = os.Args
	if len(args) <= 1 {
		fmt.Print("usage: gcat.exe filename\n")
		log.Fatal("No argument filename given.")
		os.Exit(1)
	}

	const maxLines int = 40
	file, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var screenCount int = 0

	for index, line := range lines {
		fmt.Printf("%d:     %s\n", index+1, line)
		screenCount += 1
		if screenCount%maxLines == 0 {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("\n\n---------------- Press ENTER to continue: ")
			text, _ := reader.ReadString('\n')
			fmt.Print("\n\n")
			unused(text)
		}
	}
	fmt.Printf("\n%d lines.\n", len(lines))
	file.Close()
}
