package rabbitmq

import (
	"encoding/json"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/model"

	"github.com/streadway/amqp"
)

// ViolinSchdulerToXXX : P
func ViolinSchdulerToXXX(action model.Control) error {
	qCreate, err := Channel.QueueDeclare(
		"to_xxx",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		logger.Logger.Println("ViolinSchdulerToXXX: Failed to declare a create queue")
		return err
	}

	body, _ := json.Marshal(action)
	err = Channel.Publish(
		"",
		qCreate.Name,
		false,
		false,
		amqp.Publishing{
			ContentType:     "text/plain",
			ContentEncoding: "utf-8",
			Body:            body,
		})
	if err != nil {
		logger.Logger.Println("ViolinSchdulerToXXX: Failed to register publisher")
		return err
	}

	return nil
}
