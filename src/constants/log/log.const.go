package constants_log

const (
	FILE_NAME_ERROR_GENERAL  = "log_err_general"
	FILE_NAME_ERROR_DATABASE = "log_err_database"

	OPERATION_DATABASE      = "database"
	OPERATION_MAILS_REQUEST = "mails request"

	ERROR_DATA_BASE_REQUEST        = "There was an error in the database request"
	ERROR_DATA_BASE_CREATE_REQUEST = "There was an error creating the request"

	ERROR_CREATE_LOG     = "log file could not be created"
	ERROR_CREATE_LOGBOOK = "logbook file could not be created"
	ERROR_JSON_PARSE     = "could not parse to json"

	ERROR_FROM_MAX_IS_NOT_NUMBER = "from or max must be numbers"
	ERROR_VALUE_LESS_ZERO        = "values less than 0 are not allowed"
	ERROR_INVALID_PARAMS         = "invalid parameters entered"
)
