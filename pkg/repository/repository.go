package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.com/zharzhanov/myguru/model"
)

const (
	coursesTable = "courses_ru"
	mentorsTable = "mentors"

)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GetUser(user model.User) (string, error)
}

type Courses interface {
	GetCourses() ([]model.Course, error)
	GetTrendingCourses() ([]model.Course, error)
	GetComingSoonCourses() ([]model.Course, error)
	GetCourseByID(id int) (model.Course, error)
}

type Mentors interface {
	GetMentors() ([]model.Mentor, error)
	GetMentorByID(id int) (model.Mentor, error)
}

type Categories interface {
	GetCategories(lang string) ([]model.Category, error)
}

type Repository struct {
	Authorization
	Courses
	Mentors
	Categories
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		Categories: NewCategoriesRepository(db),
		Courses: NewCoursesRepository(db),
		Mentors: NewMentorsRepository(db),
	}
}