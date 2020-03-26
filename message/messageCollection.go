package message

import "errors"

// Messages contains all current messages
var Messages []Message

// Send will append the message to the message array
func Send(msg Message) error {
	Messages = append(Messages, msg)

	return nil
}

// Delete will delete from message array based on inputs
func Delete(msg Message) error {
	for i, element := range Messages {
		if element.Content == msg.Content {
			Messages = append(Messages[:i], Messages[i+1:]...)
			break
		}
	}

	return errors.New("Message does not exist")
}

// Get returns a single message from the messages structure
func Get(msg Message) (Message, error) {
	for _, element := range Messages {
		if element.Content == msg.Content {
			return element, nil
		}
	}

	return Message{}, errors.New("Message does not exist")
}

// GetMine returns all of the messages that have been sent by a user
func GetMine() ([]Message, error) {

	return Messages, nil
}
