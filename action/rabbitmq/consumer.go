package rabbitmq

import (
	"encoding/json"
	"hcc/violin-scheduler/lib/logger"
	"hcc/violin-scheduler/model"
)

// ViolinToViolinSchedueler : Consume Viola command
func ViolinToViolinSchedueler() error {
	qCreate, err := Channel.QueueDeclare(
		"to_violin_scheduler",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		logger.Logger.Println("ViolinToViolinSchedueler: Failed to get to_violin_scheduler")
		return err
	}

	msgsCreate, err := Channel.Consume(
		qCreate.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Logger.Println("ViolinToViolinSchedueler: Failed to register to_violin_scheduler")
		return err
	}

	go func() {
		for d := range msgsCreate {
			logger.Logger.Printf("ViolinToViolinSchedueler: Received a create message: %s\n", d.Body)

			var control model.Control
			err = json.Unmarshal(d.Body, &control)
			if err != nil {
				logger.Logger.Println("ViolinToViolinSchedueler: Failed to unmarshal to_violin_scheduler data")
				// return
			}
			logger.Logger.Println("[ViolinToViolinSchedueler]RabbitmQ Receive: ", control)
			//To-Do******************************/
			//
			//*************************** */
			//

		}
	}()

	return nil
}
