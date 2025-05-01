package repo

import (
	"time"

	"github.com/vnFuhung2903/rubyams-backend/internal/entity"
	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll() ([]*entity.Course, error)
	FindByCourseName(courseName string) ([]*entity.Course, error)
	FindByID(id uint) (*entity.Course, error)
	CreateCourse(name string, teacher string, semester float32, Weekday time.Weekday, startTime time.Time, endTime time.Time, capacity uint) (*entity.Course, error)
	UpdateCourse(course *entity.Course, field string, data interface{}) error
}

type courseRepository struct {
	Db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{Db: db}
}

func (cr *courseRepository) FindAll() ([]*entity.Course, error) {
	var courses []*entity.Course
	res := cr.Db.Find(&courses)
	if res.Error != nil {
		return nil, res.Error
	}
	return courses, nil
}

func (cr *courseRepository) FindByCourseName(courseName string) ([]*entity.Course, error) {
	var courses []*entity.Course
	res := cr.Db.Find(&courses, entity.Course{CourseName: courseName})
	if res.Error != nil {
		return nil, res.Error
	}
	return courses, nil
}

func (cr *courseRepository) FindByID(id uint) (*entity.Course, error) {
	var course entity.Course
	res := cr.Db.First(&course, entity.Course{ID: id})
	if res.Error != nil {
		return nil, res.Error
	}
	return &course, nil
}

func (cr *courseRepository) CreateCourse(name string, teacher string, semester float32, Weekday time.Weekday, startTime time.Time, endTime time.Time, capacity uint) (*entity.Course, error) {
	newCourse := &entity.Course{
		CourseName:  name,
		TeacherName: teacher,
		Semester:    semester,
		Weekday:     Weekday,
		InProgress:  false,
		StartTime:   startTime,
		EndTime:     endTime,
		Capacity:    capacity,
	}
	res := cr.Db.Create(newCourse)
	if res.Error != nil {
		return nil, res.Error
	}
	return newCourse, nil
}

func (cr *courseRepository) UpdateCourse(course *entity.Course, field string, data interface{}) error {
	res := cr.Db.Model(course).Update(field, data)
	return res.Error
}
