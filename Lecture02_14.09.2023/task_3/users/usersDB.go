package users

type User struct {
	Name        string
	PhoneNumber string
	// Some other info
}

// NewUser TODO: Create user registration and validation
func NewUser(name, phoneNumber string) *User {
	return &User{
		Name:        name,
		PhoneNumber: phoneNumber,
	}
}
