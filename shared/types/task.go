package types

import "time"

type Task struct {
	ID uint `json:"id"`

	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`

	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
}
