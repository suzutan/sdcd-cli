package model

import "time"

type Event struct {
	ID           int        `json:"id"`
	PipelineID   int        `json:"pipelineId"`
	SHA          string     `json:"sha"`
	Type         string     `json:"type"`
	Status       string     `json:"status"`
	CreateTime   *time.Time `json:"createTime"`
	Creator      EventCreator `json:"creator"`
	ParentEventID *int      `json:"parentEventId"`
}

type EventCreator struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}
