package domain

import "sort"

type MessageList []Message

// SortByCreatedAtDescending orders messages by age
func (ml *MessageList) SortByCreatedAtDescending() {
	sort.Slice(*ml, func(i, j int) bool {
		return (*ml)[i].CreatedAt.After((*ml)[j].CreatedAt)
	})
}

// AddMessage adds a new message to the list
func (ml *MessageList) AddMessage(msg *Message) {
	if msg != nil {
		*ml = append(*ml, *msg)
	}
}

// GetAllMessages returns all messages as a slice of []Message
func (ml *MessageList) GetAllMessages() []Message {
	return *ml
}
