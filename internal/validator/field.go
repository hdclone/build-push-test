package validator

import (
	"strings"
)

type Field struct {
	messages []Message
}

func (m Field) Error() string {
	var texts []string
	for _, message := range m.messages {
		texts = append(texts, string(message))
	}
	return strings.Join(texts, ",")
}

func (m *Field) Append(messages ...Message) {
	m.messages = append(m.messages, messages...)
}

func (m *Field) Empty() bool {
	return len(m.messages) == 0
}
