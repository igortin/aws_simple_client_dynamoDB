package main

import "encoding/json"

// Event structure
type Event struct {
	EventID   string
	EventInfo string
	Body      json.RawMessage
}
