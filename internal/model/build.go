package model

import "time"

type Build struct {
	ID         int        `json:"id"`
	JobID      int        `json:"jobId"`
	EventID    int        `json:"eventId"`
	Status     string     `json:"status"`
	SHA        string     `json:"sha"`
	CreateTime *time.Time `json:"createTime"`
	StartTime  *time.Time `json:"startTime"`
	EndTime    *time.Time `json:"endTime"`
	Number     float64    `json:"number"`
}

type Step struct {
	Name      string     `json:"name"`
	Code      *int       `json:"code"`
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
}

type LogLine struct {
	T int64  `json:"t"`
	M string `json:"m"`
	N int    `json:"n"`
}

type LogPage struct {
	Lines    []LogLine
	NextPage int
}
