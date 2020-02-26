package controller

import (
	"github.com/LostLaser/go-twitter-rpc/user"
)

// AddUser adds a new user to the collection of users
func (t *Task) AddUser(u user.User, reply *user.User) error {
	user.Add(u)

	return nil
}

// AddFollower appends a new entry to the specified users' followers list
func (t *Task) AddFollower(u *user.User, reply *user.User) error {
	// err := user.AddFollower(u, follower)

	return nil
}

// RemoveFollowerByID removes a user by ID
func (t *Task) RemoveFollowerByID(u *user.User, reply *user.User) error {
	// user.RemoveFollowerByID(u, follower)

	return nil
}
