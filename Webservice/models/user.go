package models

import (
	"errors"
	"fmt"
)

//User type
type User struct {
	ID        int
	Firstname string
	Lastname  string
}

var (
	users  []*User
	nextID = 1
)

//GetUsers returns all user
func GetUsers() []*User {
	return users
}

//AddUser adds a new user
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New user must not include ID")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

//GetUserByID gets a user using id
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

//UpdateUser updates a user
func UpdateUser(u User) (User, error) {
	for i, usr := range users {
		if usr.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

//RemoveUserByID removes a user
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
