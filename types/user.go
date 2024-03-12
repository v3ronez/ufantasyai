package types

const UserContextKey = "user"

type AuthenticateUser struct {
	Email    string
	LoggedIn bool
}
