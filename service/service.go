package service

import "enterprise-hello-world/messsages"

type MessageService interface {
	NewCommonMessage() messsages.CommonMessage
}
