package repo

import (
	"time"

	"github.com/vnFuhung2903/rubyams-backend/internal/entity"
	"gorm.io/gorm"
)

type StudentRepository interface {
	FindAll() ([]*entity.Student, error)
	FindByNFTAddress(nftAddress string) (*entity.Student, error)
	FindByStudentNumber(studentNumber uint) (*entity.Student, error)
	CreateStudent(userID uint, name string, nftAddress string, studentNumber uint, academicYear uint, enrolledAt time.Time) (*entity.Student, error)
	UpdateStudent(student *entity.Student, field string, data interface{}) error
}

type studentRepository struct {
	Db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{Db: db}
}

func (sr *studentRepository) FindAll() ([]*entity.Student, error) {
	var students []*entity.Student
	res := sr.Db.Find(&students)
	if res.Error != nil {
		return nil, res.Error
	}
	return students, nil
}

func (sr *studentRepository) FindByNFTAddress(nftAddress string) (*entity.Student, error) {
	var student entity.Student
	res := sr.Db.First(&student, entity.Student{NFTAddress: nftAddress})
	if res.Error != nil {
		return nil, res.Error
	}
	return &student, nil
}

func (sr *studentRepository) FindByStudentNumber(studentNumber uint) (*entity.Student, error) {
	var student entity.Student
	res := sr.Db.First(&student, entity.Student{StudentNumber: studentNumber})
	if res.Error != nil {
		return nil, res.Error
	}
	return &student, nil
}

func (sr *studentRepository) CreateStudent(userID uint, name string, nftAddress string, studentNumber uint, academicYear uint, enrolledAt time.Time) (*entity.Student, error) {
	newStudent := &entity.Student{
		UserID:        userID,
		Name:          name,
		NFTAddress:    nftAddress,
		StudentNumber: studentNumber,
		AcademicYear:  academicYear,
		EnrolledAt:    enrolledAt,
		Graduated:     false,
		Wallet:        "",
	}
	res := sr.Db.Create(newStudent)
	if res.Error != nil {
		return nil, res.Error
	}
	return newStudent, nil
}

func (sr *studentRepository) UpdateStudent(student *entity.Student, field string, data interface{}) error {
	res := sr.Db.Model(student).Update(field, data)
	return res.Error
}
