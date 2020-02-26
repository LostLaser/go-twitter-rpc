package user

import "errors"

// Users is the collection of all signed up users
var Users = make(map[int]User)

// Add adds a new user to the collection of users
func Add(usr User) error {
	if _, ok := Users[usr.ID]; !ok {
		Users[usr.ID] = usr
	} else {
		return errors.New("User already exists")
	}

	return nil
}
