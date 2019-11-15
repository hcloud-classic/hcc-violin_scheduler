package init

import (
	"hcc/violin-scheduler/action/rabbitmq"
	"hcc/violin-scheduler/lib/logger"
)

func rabbitmqInit() error {
	err := rabbitmq.PrepareChannel()
	if err != nil {
		return err
	}

	go func() {
		forever := make(chan bool)
		logger.Logger.Println("RabbitMQ forever channel ready.")
		<-forever
	}()

	return nil
}
