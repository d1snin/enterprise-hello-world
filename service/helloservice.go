package service

import (
	"enterprise-hello-world/config"
	"enterprise-hello-world/messsages"
	"fmt"
)

type HelloMessageService struct {
	conf *config.Configuration
}

func NewHelloMessageService(conf *config.Configuration) *HelloMessageService {
	return &HelloMessageService{conf: conf}
}

func (s *HelloMessageService) NewCommonMessage() messsages.CommonMessage {
	return messsages.CommonMessage{
		Message: fmt.Sprintf("Hello, %v !", s.conf.MessageReceiver),
	}
}
