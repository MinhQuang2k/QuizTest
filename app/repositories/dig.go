package repositories

import (
	"go.uber.org/dig"
)

// Inject repositories
func Inject(container *dig.Container) error {
	_ = container.Provide(NewCategoryRepository)
	_ = container.Provide(NewExamRepository)
	_ = container.Provide(NewGroupQuestionRepository)
	_ = container.Provide(NewQuestionRepository)
	_ = container.Provide(NewRoomRepository)
	_ = container.Provide(NewSubjectRepository)
	_ = container.Provide(NewUserRepository)
	return nil
}
