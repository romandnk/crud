package task

import "time"

type Task struct {
	Id           int       `json:"id"`
	CreationTime time.Time `json:"creation_time"`
	UpdatingTime time.Time `json:"updating_time"`
	Message      string    `json:"message"`
}
