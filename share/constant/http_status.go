package constant

const (
	DEFAULT_ERROR = iota + 40000
	INVALID_PARAMS
	INVALID_QUERY
	INVALID_LOGIN
	INVALID_TOKEN
	INVALID_BODY
	PERMISSION_DENY
	RESOURCE_EXIST
	REQUEST_INVALIDATED
	REQUESTED_RESOURCE_NOT_FOUND
	TOO_MANY_REQUESTS
	MISSION_EXECUTION_FAILED
)
const (
	SUCCESS = iota + 20000
	CREATED_SUCCESS
)
