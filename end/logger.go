package end

import "hcc/violin-scheduler/lib/logger"

func loggerEnd() {
	_ = logger.FpLog.Close()
}
