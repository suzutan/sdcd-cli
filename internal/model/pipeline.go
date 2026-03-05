package model

import "time"

type Pipeline struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ScmURI      string     `json:"scmUri"`
	ScmContext  string     `json:"scmContext"`
	CreateTime  *time.Time `json:"createTime"`
	LastEventID int        `json:"lastEventId"`
	State       string     `json:"state"`
}
