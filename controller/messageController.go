package controller

import (
	"github.com/LostLaser/go-twitter-rpc/message"
)

// SendMessage will append the message to the message array
func (t *Task) SendMessage(msg message.Message, reply *message.Message) error {
	err := message.Send(msg)

	return err
}

// DeleteMessage will delete from message array based on inputs
func (t *Task) DeleteMessage(msg message.Message, reply *message.Message) error {
	err := message.Delete(msg)

	return err
}

// GetMessage returns a message from the messages structure
func (t *Task) GetMessage(msg message.Message, reply *message.Message) error {
	val, err := message.Get(msg)

	*reply = val

	return err
}

// GetMyMessages returns all of the messages that have been sent by a user
func (t *Task) GetMyMessages(msg message.Message, reply *[]message.Message) error {
	val, err := message.GetMine()

	*reply = val

	return err
}
