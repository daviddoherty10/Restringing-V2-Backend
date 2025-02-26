package entity

import (
	"time"
)

type LoggingMiddleware struct {
	Id              uint
	StatusCode      int
	ClientIP        string
	UserAgent       string
	Referer         string
	RequestMethod   string
	RequestURL      string
	Hearders        string
	ResponseBodyStr string
	Duration        time.Duration
	CreatedAt       time.Time
}
