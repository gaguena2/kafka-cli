package main

import (
	"fmt"

	"github.com/gaguena2/kafka-cli/app/cmd"
)

func main() {
	err := cmd.CreateCommand()

	if err != nil {
		fmt.Println(err)
	}
}
