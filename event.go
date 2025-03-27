package wlib

import "time"

// Event is a Message for MQ
type Event struct {
	Timestamp time.Time
	Type      string
	Ev        Message
}
