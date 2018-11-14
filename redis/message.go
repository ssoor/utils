package redis

// Message represents a message notification.
type Message struct {
	// The originating channel.
	Channel string

	// The matched pattern, if any
	Pattern string

	// The message data.
	Data []byte
}
