package domain

import "sort"

type MessageList []Message

// SortByCreatedAtDescending orders messages by age
func (ml *MessageList) SortByCreatedAtDescending() {
	sort.Slice(*ml, func(i, j int) bool {
		return (*ml)[i].CreatedAt().After((*ml)[j].CreatedAt())
	})
}

// GetContents returns a slice of all message contents
func (ml *MessageList) GetContents() []string {
	contents := make([]string, len(*ml))
	for i, msg := range *ml {
		contents[i] = msg.Content() + "- said @" + msg.UserName()
	}
	return contents
}

func (ml *MessageList) AddMessage(msg *Message) {
	if msg != nil {
		*ml = append(*ml, *msg)
	}
}
