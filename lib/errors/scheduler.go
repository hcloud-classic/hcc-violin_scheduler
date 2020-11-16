package errors

const (
	ShcedulerInternalInitFail                 = violinScheduler + internal + initFail
	ShcedulerInternalOperationFail            = violinScheduler + internal + operationFail
	ShcedulerInternalTimeStampConversionError = violinScheduler + internal + timestampConversionError

	ShcedulerHandlerFaild = violinScheduler + schedulehandlerError

	ShcedulerGrpcArgumentError = violinScheduler + grpc + argumentError
	ShcedulerGrpcRequestError  = violinScheduler + grpc + requestError
	ShcedulerGrpcGetNodesError = violinScheduler + grpc + getNodesError

	ShcedulerSQLOperationFail = violinScheduler + sql + operationFail
	ShcedulerSQLNoResult      = violinScheduler + sql + noResult
	ShcedulerSQLArgumentError = violinScheduler + sql + argumentError
)
