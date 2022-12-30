package repository

import (
	"cornjobmailer/features/user/domain"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type repoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.Repository {
	return &repoQuery{
		db: db,
	}
}

// Add implements domain.Repository
func (rq *repoQuery) Add(data domain.UserCore) (domain.UserCore, error) {
	var cnv User = FromDomain(data)

	if err := rq.db.Create(&cnv).Error; err != nil {
		log.Error("error on adding user", err.Error())
		return domain.UserCore{}, err
	}
	return ToDomain(cnv), nil
}

// Login implements domain.Repository
func (rq *repoQuery) Login(input domain.UserCore) (domain.UserCore, error) {
	var cnv User
	if err := rq.db.Table("users").First(&cnv, "email = ?", input.Email).Error; err != nil {
		log.Error("error on get user login", err.Error())
		return domain.UserCore{}, nil
	}
	res := ToDomain(cnv)
	return res, nil
}

// GetAll implements domain.Repository
func (rq *repoQuery) GetAll() ([]domain.UserCore, error) {
	var data []User
	if err := rq.db.Find(&data).Error; err != nil {
		log.Error("error on query get all data user", err.Error())
		return nil, err
	}
	return ToDomainArray(data), nil
}

// Edit implements domain.Repository
func (rq *repoQuery) Edit(id uint, email string) (domain.UserCore, error) {
	panic("unimplemented")
}
