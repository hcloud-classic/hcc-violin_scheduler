package end

import "hcc/violin-scheduler/action/grpc/client"

// MainEnd : Main ending function
func MainEnd() {
	// mysqlEnd()
	loggerEnd()
	client.End()
}
