package apis

import (
	"strconv"

	"gopkg.in/tucnak/telebot.v2"
)

// Request is an interface for all Telegram request types (message, callback, etc.).
type Request interface {
	ReqID() string
	Sender() *telebot.User
}

// Message wraps Telebot's Message type.
type Message struct {
	*telebot.Message
}

// ReqID returns the ID of m.
func (m *Message) ReqID() string {
	return strconv.Itoa(m.ID)
}

// Sender returns the sender of m.
func (m *Message) Sender() *telebot.User {
	return m.Message.Sender
}

// Callback wraps Telebot's Callback type.
type Callback struct {
	*telebot.Callback
}

// ReqID returns the ID of cb.
func (cb *Callback) ReqID() string {
	return cb.ID
}

// Sender returns the sender of cb.
func (cb *Callback) Sender() *telebot.User {
	return cb.Callback.Sender
}
