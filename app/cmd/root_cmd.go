package cmd

import (
	"errors"

	flag "github.com/spf13/pflag"
)

func Execute() error {
	var topic = flag.StringP("topic", "t", "", "Topic Name")
	var message = flag.StringP("message", "m", "", "Message send")
	flag.Parse()

	if *topic == "" {
		return errors.New("invalid parameter command line: \"topic name\"")
	}
	if *message == "" {
		return errors.New("invalid parameter command line: \"message\"")
	}
	println(*topic)
	return nil
}
