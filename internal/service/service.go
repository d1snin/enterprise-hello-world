package service

import (
	"enterprise-hello-world/internal/messages"
)

// MessageService is a service that contains a logic for creating messages.CommonMessage instances.
type MessageService interface {
	NewCommonMessage() messages.CommonMessage
}
