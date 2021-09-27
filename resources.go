package main

import "time"

type timeResource struct {
	CurrentTime time.Time `json:"current_time"`
}
