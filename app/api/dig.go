package api

import (
	"go.uber.org/dig"
)

// Inject apis
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryAPI)
	_ = container.Provide(NewExamAPI)
	_ = container.Provide(NewGroupQuestionAPI)
	_ = container.Provide(NewQuestionAPI)
	_ = container.Provide(NewRoomAPI)
	_ = container.Provide(NewSubjectAPI)
	_ = container.Provide(NewUserAPI)
	return nil
}
