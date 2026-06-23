package main

type User struct {
	Name string
	Age  int
}

func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u User) GetName() string {
	return u.Name
}
