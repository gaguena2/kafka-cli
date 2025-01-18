package cmd

import (
	"fmt"

	"github.com/gaguena2/kafka-cli/internal/usecase"
)

func CommandRun(usecase usecase.UseCase) {
	err := usecase.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
