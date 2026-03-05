package model

import "time"

type Secret struct {
	ID         int        `json:"id"`
	PipelineID int        `json:"pipelineId"`
	Name       string     `json:"name"`
	AllowInPR  bool       `json:"allowInPR"`
	CreateTime *time.Time `json:"createTime"`
}
