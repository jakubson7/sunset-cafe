package models

import "time"

type Timestamp struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func (m *Timestamp) Validate() error {
	e := newValidationError("Timestamp")

	if m.CreatedAt == 0 {
		return e.From("CreatedAt can't be 0")
	} else if m.CreatedAt < 0 {
		return e.From("CreatedAt can't be smaller than 0")
	}
	if m.UpdatedAt == 0 {
		return e.From("UpdatedAt can't be 0")
	} else if m.UpdatedAt < 0 {
		return e.From("UpdatedAt can't be smaller than 0")
	}

	return nil
}

func NewTimestamp() Timestamp {
	now := time.Now().Unix()
	return Timestamp{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
