package model

// User - struct for create, write, update , delete user.
type User struct {
	Email string `toml:"email"`
}

func NewUser(email string) *User {
	return &User{
		Email: email,
	}
}
