package http

// The error values below consit of two lowercased words because API errors
// should always be short and sweet.
const (
	AccountLocked       = "account locked"
	ExpiredSession      = "expired session"
	InvalidCreds        = "invalid credentials"
	InvalidEmail        = "invalid email"
	InvalidFile         = "invalid file"
	InvalidID           = "invalid id"
	InvalidInput        = "invalid input"
	InvalidJSON         = "invalid json"
	InvalidJSONP        = "invalid jsonp"
	InvalidXML          = "invalid xml"
	InvalidYAML         = "invalid yaml"
	InvalidParameter    = "invalid parameter"
	InvalidPassword     = "invalid password"
	InvalidPermission   = "invalid permission"
	InvalidRegistration = "invalid registration"
	InvalidSession      = "invalid session"
	InvalidUser         = "invalid user"
	MissingSession      = "missing session"
	MissingUser         = "missing user"
)

// ErrorMap is a map of error messages to HTTP status codes.
//
// These error messages are used by Response to match standard error messages
// with thier proper HTTP status codes.
var ErrorMap = map[string]int{
	AccountLocked:       StatusForbidden,
	ExpiredSession:      StatusBadRequest,
	InvalidCreds:        StatusBadRequest,
	InvalidEmail:        StatusBadRequest,
	InvalidFile:         StatusBadRequest,
	InvalidID:           StatusBadRequest,
	InvalidInput:        StatusBadRequest,
	InvalidJSON:         StatusBadRequest,
	InvalidJSONP:        StatusBadRequest,
	InvalidXML:          StatusBadRequest,
	InvalidYAML:         StatusBadRequest,
	InvalidParameter:    StatusBadRequest,
	InvalidPassword:     StatusBadRequest,
	InvalidPermission:   StatusForbidden,
	InvalidRegistration: StatusBadRequest,
	InvalidSession:      StatusBadRequest,
	InvalidUser:         StatusBadRequest,
	MissingSession:      StatusBadRequest,
	MissingUser:         StatusBadRequest,
}

// ApplyErrorCode inserts an error key and status code value into the ErrorMap.
func ApplyErrorCode(err error, statusCode int) bool {
	message := err.Error()
	_, exists := ErrorMap[message]
	if exists {
		return false
	}
	ErrorMap[message] = statusCode
	return true
}
