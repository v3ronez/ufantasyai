package types

import (
	"time"

	"github.com/google/uuid"
)

type ImageStatus int //enum in golang

const (
	ImageStatusFailed ImageStatus = iota
	ImageStatusPending
	ImageStatusCompleted
)

type Image struct {
	ID            int `bun:"id,pk,autoincrement"`
	UserId        uuid.UUID
	Prompt        string
	Status        ImageStatus
	ImageLocation string
	Deleted       bool `bun:"default:'false'"`
	CreatedAt     time.Time
	DeletedAt     time.Time
}
