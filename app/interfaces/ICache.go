package interfaces

import "quiztest/pkg/redis"

// ICache interface
type ICache interface {
	GetInstance() redis.IRedis
}
