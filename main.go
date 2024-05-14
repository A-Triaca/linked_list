package main

import (
	"bufio"
	"fmt"
	"linked_list/list"
	"os"
	"strconv"
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
		value, err := strconv.ParseInt(line, 10, 32)

		if err != nil {
			fmt.Print("Only integers supported")
		}

		l.Insert(value)
		fmt.Printf(l.String())
	}
}
