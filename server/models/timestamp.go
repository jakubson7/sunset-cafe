package models

import (
	"errors"
	"fmt"
	"time"
)

type Timestamp struct {
	CreatedAt int64 `json:"createdAt"`
	UpdatedAt int64 `json:"updatedAt"`
}

func (m *Timestamp) Validate() error {
	if m.CreatedAt == 0 {
		return m.efrom("CreatedAt can't be 0")
	} else if m.CreatedAt < 0 {
		return m.efrom("CreatedAt can't be smaller than 0")
	}
	if m.UpdatedAt == 0 {
		return m.efrom("UpdatedAt can't be 0")
	} else if m.UpdatedAt < 0 {
		return m.efrom("UpdatedAt can't be smaller than 0")
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

func (m *Timestamp) efrom(text string) error {
	return errors.New(fmt.Sprintf("(Timestamp) -> %s", text))
}
