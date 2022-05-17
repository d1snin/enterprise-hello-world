package service

import (
	"enterprise-hello-world/internal/config"
	"enterprise-hello-world/internal/messages"
	"fmt"
)

// HelloMessageService is a MessageService implementation that contains the base logic.
type HelloMessageService struct {
	conf *config.Configuration
}

// NewHelloMessageService creates a new HelloMessageService instance.
func NewHelloMessageService(conf *config.Configuration) *HelloMessageService {
	return &HelloMessageService{conf: conf}
}

// NewCommonMessage creates a new messages.CommonMessage instance that contains a greeting.
func (s *HelloMessageService) NewCommonMessage() messages.CommonMessage {
	return messages.CommonMessage{
		Message: fmt.Sprintf("Hello, %v!", s.conf.MessageReceiver),
	}
}
