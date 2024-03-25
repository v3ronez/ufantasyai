package types

import (
	"time"

	"github.com/google/uuid"
)

type ImageStatus int //enum in golang

const (
	ImageStatusFailed ImageStatus = iota
	ImageStatusPending
	ImageStatusComplted
)

type Image struct {
	ID        int `bun:"id, pk, autoincrement"`
	UserId    uuid.UUID
	Status    ImageStatus
	CreatedAt time.Time
}
