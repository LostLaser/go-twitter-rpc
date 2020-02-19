package message

import "time"

// Message contains information for one individual message
type Message struct {
	Username  string
	Content   string
	TimeStamp time.Time
}
