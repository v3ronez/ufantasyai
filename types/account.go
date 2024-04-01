package types

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        int       `bun:"id,pk,unique,autoincrement" json:"id"`
	UserId    uuid.UUID `bun:"user_id,unique,type:uuid" json:"user_id"`
	UserName  string    `bun:"username" json:"username"`
	Credits   int       `bun:"credits" json:"credits"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
}
