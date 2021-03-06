package shuftipro

import "strings"

// errorField represents the error field in the response when it isn't empty.
type errorField struct {
	Error Error `json:"error"`
}

// Error represents an error.
type Error struct {
	Service string `json:"service"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

// Error implements the error interface for the Error.
func (e Error) Error() string {
	b := strings.Builder{}

	if len(e.Service) > 0 {
		b.WriteString("service: '")
		b.WriteString(e.Service)
		b.WriteByte('\'')
	}
	if len(e.Key) > 0 {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		b.WriteString("key: '")
		b.WriteString(e.Key)
		b.WriteByte('\'')
	}
	if len(e.Message) > 0 {
		if b.Len() > 0 {
			b.WriteString(" | ")
		}
		b.WriteString("message: '")
		b.WriteString(e.Message)
		b.WriteByte('\'')
	}

	return b.String()
}
