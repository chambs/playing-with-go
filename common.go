package main

// User ...
type User struct {
	ID    int32
	name  string
	Email string
}

func (u *User) setname(name string) {
	u.name = name
}
func (u *User) setEmail(email string) {
	u.Email = email
}

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func div(a, b int) int {
	return a / b
}
