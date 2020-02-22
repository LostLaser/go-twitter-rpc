package user

import "errors"

// User contains all of the information relevant to an individual user
type User struct {
	FirstName string
	LastName  string
	ID        int
	NickName  string
	Password  string
	Followers map[int]*User
}

// AddFollower appends a new entry to the users' followers list
func AddFollower(usr *User, follower *User) error {
	if _, found := usr.Followers[follower.ID]; !found {
		usr.Followers[follower.ID] = follower
	} else {
		return errors.New("User is already in follower list")
	}

	return nil
}

// RemoveFollowerByID removes a user by ID
func RemoveFollowerByID(usr *User, follower *User) error {
	if _, found := usr.Followers[follower.ID]; found {
		delete(usr.Followers, follower.ID)
	} else {
		return errors.New("User does not exist in follower list with ID: " + string(follower.ID))
	}

	return nil
}
