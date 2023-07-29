package models

import "time"

type Timestamp struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func NewTimestamp() Timestamp {
	now := time.Now().Unix()
	return Timestamp{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
