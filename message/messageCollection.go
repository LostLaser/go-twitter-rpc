package message

import "errors"

// Messages contains all current messages
var Messages []Message

// MsgTask is a method type for RPC
type MsgTask int

// SendMessage will append the message to the message array
func SendMessage(msg Message) error {
	Messages = append(Messages, msg)

	return nil
}

// DeleteMessage will delete from message array based on inputs
func DeleteMessage(msg Message) error {
	for i, element := range Messages {
		if element.Content == msg.Content {
			Messages = append(Messages[:i], Messages[i+1:]...)
			break
		}
	}

	return errors.New("Message does not exist")
}

// GetMessage returns a message from the messages structure
func GetMessage(msg Message) (Message, error) {
	for _, element := range Messages {
		if element.Content == msg.Content {
			return element, nil
		}
	}

	return Message{}, errors.New("Message does not exist")
}

// GetMyMessages returns all of the messages that have been sent by a user
func GetMyMessages() ([]Message, error) {

	return Messages, nil
}
