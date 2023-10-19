package mail

import (
	"go.uber.org/dig"
)

// Inject dbs
func Inject(container *dig.Container) error {
	_ = container.Provide(NewMail)
	return nil
}
