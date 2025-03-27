package wlib

import "time"

type Message any

type MessageOptions struct {
	SubTopic string
}

type MessageOp func(*MessageOptions)

type MQSink interface {
	Send(Message, ...MessageOp) SendToken
}

// SendToken is a Future on the successful sending of a Message through the Sink
// interface.
type SendToken interface {
	// Wait blocks indefinately for the Send operation to complete.  The result
	// indicates successful sending of the Message.
	Wait() bool

	// WaitTimeout blocks for up to the timeout for the Send operation to
	// complete.  The result indicates successful sending of the Message.  A
	// negative result does not indicate that the Message _may_ be sent
	// successfully in the future.
	WaitTimeout(time.Duration) bool

	// Done returns a channel which is closed upon successful sending or
	// cancelation of the Message.
	Done() <-chan struct{}

	// Error returns nil if the Message is either pending being sent, or was
	// successfully sent.  Non-nil responses indicate the cause of the sending
	// failure.
	Error() error
}

func WithSubTopic(t string) MessageOp {
	return func(mo *MessageOptions) {
		mo.SubTopic = t
	}
}
