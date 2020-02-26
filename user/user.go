package user

import (
	"errors"

	"github.com/LostLaser/go-twitter-rpc/message"
)

// User contains all of the information relevant to an individual user
type User struct {
	FirstName    string
	LastName     string
	ID           int
	NickName     string
	Password     string
	Following    map[int]*User
	SentMessages []message.Message
}

// AddFollower appends a new entry to the users' following list
func AddFollower(usr *User, target *User) error {
	if _, found := usr.Following[target.ID]; !found {
		usr.Following[target.ID] = target
	} else {
		return errors.New("User is already in following list")
	}

	return nil
}

// RemoveFollowerByID removes a user by ID
func RemoveFollowerByID(usr *User, target *User) error {
	if _, found := usr.Following[target.ID]; found {
		delete(usr.Following, target.ID)
	} else {
		return errors.New("User does not exist in following list with ID: " + string(target.ID))
	}

	return nil
}
