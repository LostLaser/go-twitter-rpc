package user

import "errors"

// Users is the collection of all signed up users
var Users = make(map[int]User)

// CollectionAdd adds a new user to the collection of users
func CollectionAdd(usr User) error {
	if _, exist := Users[usr.ID]; exist {
		return errors.New("User already exists")
	}

	Users[usr.ID] = usr
	return nil
}

// CollectionDelete removes a user from the collection of users
func CollectionDelete(usr User) error {
	if _, exist := Users[usr.ID]; exist {
		delete(Users, usr.ID)
	} else {
		return errors.New("User does not exist")
	}

	return nil
}
