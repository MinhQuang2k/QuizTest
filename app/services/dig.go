package services

import (
	"go.uber.org/dig"
)

// Inject services
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryService)
	_ = container.Provide(NewExamService)
	_ = container.Provide(NewGroupQuestionService)
	_ = container.Provide(NewQuestionService)
	_ = container.Provide(NewRoomService)
	_ = container.Provide(NewSubjectService)
	_ = container.Provide(NewUserService)
	return nil
}
