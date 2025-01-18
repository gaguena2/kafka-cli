package cmd

import (
	"errors"

	"github.com/gaguena2/kafka-cli/internal/usecase"
	flag "github.com/spf13/pflag"
)

func CreateCommand() error {
	var topic = flag.StringP("topic", "t", "", "Topic Name")
	var message = flag.StringP("message", "m", "", "Message send")
	flag.Parse()

	if *topic == "" {
		return errors.New("invalid parameter command line: \"topic name\"")
	}
	if *message == "" {
		return errors.New("invalid parameter command line: \"message\"")
	}
	producerUseCase := usecase.NewProducerUseCase(topic, message)
	CommandRun(producerUseCase)
	return nil
}
