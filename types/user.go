package types

import "github.com/google/uuid"

const UserContextKey = "user"

type AuthenticateUser struct {
	ID       uuid.UUID
	Email    string
	LoggedIn bool

	Account
}
