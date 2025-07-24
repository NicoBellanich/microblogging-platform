package domain

import "sort"

type Feed []MessageList

// AddMessageList adds a new MessageList to the Feed
func (f *Feed) AddMessageList(ml *MessageList) {
	if ml != nil {
		*f = append(*f, *ml)
	}
}

// SortAllMessagesDescending orders all messages across all MessageLists by CreatedAt (descending)
func (f *Feed) SortAllMessagesDescending() {
	// Flatten all messages into a single slice
	var allMessages []Message
	for _, ml := range *f {
		allMessages = append(allMessages, ml...)
	}

	// Sort the flattened slice
	sort.Slice(allMessages, func(i, j int) bool {
		return allMessages[i].CreatedAt().After(allMessages[j].CreatedAt())
	})

	// Rebuild the Feed with sorted messages
	*f = make(Feed, 0)
	currentList := make(MessageList, 0)
	for _, msg := range allMessages {
		currentList = append(currentList, msg)
		// Optional: Start a new MessageList after a certain size (e.g., 10 messages)
		if len(currentList) >= 10 {
			*f = append(*f, currentList)
			currentList = make(MessageList, 0)
		}
	}
	if len(currentList) > 0 {
		*f = append(*f, currentList)
	}
}

// GetAllMessages returns all messages from the Feed as a single slice of Message
func (f *Feed) GetAllMessages() []Message {
	var allMessages []Message
	for _, ml := range *f {
		allMessages = append(allMessages, ml...)
	}
	return allMessages
}
