package cache

import (
	"go.uber.org/dig"
)

// Inject rdbs
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCache)
	return nil
}
