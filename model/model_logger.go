package model

import "time"

type LogRequest struct {
	Latency    time.Duration
	StatusCode int
	ClientIP   string
	Method     string
	Path       string
	Message    string
}
