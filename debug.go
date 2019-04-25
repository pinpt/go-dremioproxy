package dremioproxy

import "time"

// Debug is an interface which can handle query debug callbacks
type Debug interface {
	QueryFinished(query string, duration time.Duration, err error)
}

var debugger Debug

// SetDebug will allow you to hook into driver query results
func SetDebug(debug Debug) {
	debugger = debug
}
