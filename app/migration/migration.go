package migration

import (
	"go.uber.org/dig"

	"quiztest/app/interfaces"
	"quiztest/app/models"
)

// Migrate migrate to database
func Migrate(container *dig.Container) error {
	return container.Invoke(func(
		db interfaces.IDatabase,
	) error {
		User := models.User{}
		Category := models.Category{}
		ExamQuestion := models.ExamQuestion{}
		Exam := models.Exam{}
		GroupQuestion := models.GroupQuestion{}
		Question := models.Question{}
		Subject := models.Subject{}
		Room := models.Room{}
		Candidate := models.Candidate{}

		db.GetInstance().AutoMigrate(&User, &GroupQuestion, &Category, &Subject, &Exam, &Question, &ExamQuestion, &Room, &Candidate)

		return nil
	})
}
