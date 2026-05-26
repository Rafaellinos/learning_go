package user

import "time"

// type defines a custom type
// upper case to be available outside this package
type User struct {
	FirstName string
	LastName string
	createdAt time.Time
}

// method
func (user User) GetFullName() string {
	return user.FirstName + " " + user.LastName
}

// returns a pointer
// constructor pattern
func New(firstName, lastName string) *User {
	newUser := &User{
		FirstName: firstName,
		LastName: lastName,
		createdAt: time.Now(),
	}
	return newUser
}

