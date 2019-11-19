package init

import "hcc/violin-scheduler/lib/config"

// MainInit : Main initialization function
func MainInit() error {
	// err := syscheckInit()
	// if err != nil {
	// 	return err
	// }

	err := loggerInit()
	if err != nil {
		return err
	}

	config.Parser()

	// err = mysqlInit()
	// if err != nil {
	// 	return err
	// }

	return nil
}
