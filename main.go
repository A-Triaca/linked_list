package main

import (
	"bufio"
	"fmt"
	"linked_list/list"
	"os"
	"strconv"
	"strings"
)

const PROMPT = ">> "

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	l := list.List{}

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		input := strings.Split(line, " ")

		value, err := strconv.ParseInt(input[1], 10, 32)

		if err != nil {
			fmt.Print("Only integers supported")
		}

		if input[0] == "+" {
			l.Insert(value)
		}
		if input[0] == "-" {
			ok := l.Remove(value)
			if ok != nil {
				fmt.Printf("Error removing %d - %v\n", value, ok)
			}
		}

		fmt.Printf("Current list - %s\n", l.String())
	}
}
