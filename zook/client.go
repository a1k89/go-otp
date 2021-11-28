package zook

import "net/http"

func BadRequest(w http.ResponseWriter, message ...interface{}) {
	zook(w, 400, message...)
}

func Unathorized(w http.ResponseWriter, message ...interface{}) {
	zook(w, 401, message...)
}

func PaymentRequired(w http.ResponseWriter, message ...interface{}) {
	zook(w, 402, message...)
}

func Forbidden(w http.ResponseWriter, message ...interface{}) {
	zook(w, 403, message...)
}
func NotFound(w http.ResponseWriter, message ...interface{}) {
	zook(w, 404, message...)
}

func MethodNotAllowed(w http.ResponseWriter, message ...interface{}) {
	zook(w, 405, message...)
}

func NotAcceptable(w http.ResponseWriter, message ...interface{}) {
	zook(w, 406, message...)
}

// Conflict responds with a 409 Conflict error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func Conflict(w http.ResponseWriter, message ...interface{}) {
	zook(w, 409, message...)
}

// ResourceGone responds with a 410 Gone error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func ResourceGone(w http.ResponseWriter, message ...interface{}) {
	zook(w, 410, message...)
}

// LengthRequired responds with a 411 Length Required error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func LengthRequired(w http.ResponseWriter, message ...interface{}) {
	zook(w, 411, message...)
}

// PreconditionFailed responds with a 412 Precondition Failed error.
// Takes an optional message of either type string or type error,
// which will be returned in the response body.
func PreconditionFailed(w http.ResponseWriter, message ...interface{}) {
	zook(w, 412, message...)
}






