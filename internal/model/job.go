package model

import "time"

type Job struct {
	ID          int        `json:"id"`
	PipelineID  int        `json:"pipelineId"`
	Name        string     `json:"name"`
	State       string     `json:"state"`
	Archived    bool       `json:"archived"`
	CreateTime  *time.Time `json:"createTime"`
}
