package init

import (
	"hcc/violin-scheduler/action/grpc/client"
	"hcc/violin-scheduler/lib/config"
)

// MainInit : Main initialization function
func MainInit() error {

	err := loggerInit()
	if err != nil {
		return err
	}

	config.Parser()

	err = client.Init()
	if err != nil {
		return err
	}
	return nil
}
