package gophersauras

import "net/http"

/*
The errs package holds the error messages which have been standardized. These messages
will be used by the response to choose the proper http status code, as well as give
users of the API a set of standardized messages to which they may then use to match
to human readable messages.
*/

const (
	AccountLocked      = "account locked"
	ExpiredSession     = "expired session"
	InvalidCreds       = "invalid credentials"
	InvalidEmail       = "invalid email"
	InvalidId          = "invalid id"
	InvalidJson        = "invalid json"
	InvalidParameter   = "invalid parameter"
	InvalidPassword    = "invalid password"
	InvalidPermissions = "invalid permissions"
	InvalidSession     = "invalid session"
	MissingSession     = "missing session"
	MissingUser        = "missing user"
)

var errorMap = map[string]int{
	AccountLocked:      http.StatusInternalServerError,
	ExpiredSession:     http.StatusInternalServerError,
	InvalidCreds:       http.StatusInternalServerError,
	InvalidEmail:       http.StatusInternalServerError,
	InvalidId:          http.StatusInternalServerError,
	InvalidJson:        http.StatusInternalServerError,
	InvalidParameter:   http.StatusInternalServerError,
	InvalidPassword:    http.StatusInternalServerError,
	InvalidPermissions: http.StatusInternalServerError,
	InvalidSession:     http.StatusInternalServerError,
	MissingSession:     http.StatusInternalServerError,
	MissingUser:        http.StatusInternalServerError,
}

func ApplyErrorCode(err error, statusCode int) bool {
	message := err.Error()
	_, exists := errorMap[message]
	if exists {
		return false
	}
	errorMap[message] = statusCode
	return true
}
