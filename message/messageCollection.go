package message

// Messages contains all current messages
var Messages []Message

// Task is a method type for RPC
type Task int

// SendMessage will append the message to the message array
func (t *Task) SendMessage(msg Message, reply *Message) error {
	Messages = append(Messages, msg)

	return nil
}

// DeleteMessage will delete from message array based on inputs
func (t *Task) DeleteMessage(msg Message, reply *Message) error {
	for i, element := range Messages {
		if element.Content == msg.Content {
			Messages = append(Messages[:i], Messages[i+1:]...)
			break
		}
	}

	return nil
}

// GetMessage returns a message from the messages structure
func (t *Task) GetMessage(msg Message, reply *Message) error {
	for _, element := range Messages {
		if element.Content == msg.Content {
			*reply = element
			break
		}
	}

	return nil
}

// GetMyMessages returns all of the messages that have been sent by a user
func (t *Task) GetMyMessages(msg Message, reply *[]Message) error {
	*reply = Messages

	return nil
}
