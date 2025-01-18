package usecase

import (
	"fmt"

	"github.com/gaguena2/kafka-cli/internal/data"
)

type ProducerUseCase struct {
	Topic   *string
	Message *data.MessageData
}

func (p *ProducerUseCase) Execute() error {

	fmt.Println("Estou dentro de ProducerUseCase metodo Execute")
	fmt.Println(*p.Topic)
	return nil
}

func NewProducerUseCase(topic *string, message *string) *ProducerUseCase {
	return &ProducerUseCase{
		Topic:   topic,
		Message: nil,
	}
}
