package user

import (
	"errors"

	"github.com/LostLaser/go-twitter-rpc/server/message"
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

// AddFollower appends a new entry to a users' following list
func (usr *User) AddFollower(target *User) error {
	if _, found := usr.Following[target.ID]; !found {
		usr.Following[target.ID] = target
	} else {
		return errors.New("User is already in following list")
	}

	return nil
}

// RemoveFollowerByID removes a follower from a user object by ID
func (usr *User) RemoveFollowerByID(target *User) error {
	if _, found := usr.Following[target.ID]; found {
		delete(usr.Following, target.ID)
	} else {
		return errors.New("User does not exist in following list with ID: " + string(target.ID))
	}

	return nil
}

// GetFollowing returns all of the people that the user is currently following
func (usr *User) GetFollowing() []User {
	var following []User

	for _, follow := range usr.Following {
		following = append(following, *follow)
	}

	return following
}

// SendMessage appends a message to the calling user's SentMessages
func (usr *User) SendMessage(msg message.Message) {
	usr.SentMessages = append(usr.SentMessages, msg)
}

// GetMessages returns a list of all messages that the calling user has sent
func (usr *User) GetMessages() []message.Message {
	return usr.SentMessages
}

// GetTimeline returns a list of the messages sent by followed users
func (usr *User) GetTimeline() []message.Message {
	var timeline []message.Message

	for _, follow := range usr.Following {
		timeline = append(timeline, follow.SentMessages...)
	}

	return timeline
}
