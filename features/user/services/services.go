package services

import (
	"cornjobmailer/config"
	"cornjobmailer/features/user/domain"
	"cornjobmailer/utils/helper"
	"errors"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	qry domain.Repository
}

func New(repo domain.Repository) domain.Service {
	return &userService{
		qry: repo,
	}
}

// Register implements domain.Service
func (us *userService) Register(data domain.UserCore) (domain.UserCore, error) {

	generate, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("error on bcrypt", err.Error())
		return domain.UserCore{}, errors.New("cannot encrypt password")
	}
	data.Password = string(generate)
	orgPass := data.Password
	data.Status = "pending"
	data.Wallet.IDCurrency = 1
	data.Wallet.IDUser = data.ID
	data.Wallet.Amount = 0

	data.Mailer.Email = data.Email
	data.Mailer.IDUser = data.ID
	data.Mailer.Status = "pending"
	data.Mailer.Pin = randstr.String(20)
	err = helper.Corn(data.Name, data.Email, data.Mailer.Pin)
	if err != nil {
		log.Error("error on corn email", err.Error())
	}
	res, err := us.qry.Add(data)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate") {
			return domain.UserCore{}, errors.New("already exist")
		}
		return domain.UserCore{}, errors.New("some problem on database")
	}
	res.Password = orgPass
	return res, nil
}

// Login implements domain.Service
func (us *userService) Login(input domain.UserCore) (domain.UserCore, error) {
	res, err := us.qry.Login(input)

	if err != nil {
		log.Error(err.Error())
		if strings.Contains(err.Error(), "found") {
			return domain.UserCore{}, errors.New("Failed. Email or Password not found.")
		} else if strings.Contains(err.Error(), "table") {
			return domain.UserCore{}, errors.New("Failed. Email or Password not found.")
		}
		return domain.UserCore{}, errors.New("email not found")
	} else {
		if res.ID == 0 {
			return domain.UserCore{}, errors.New("email not found")
		}

		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(input.Password))
		if err != nil {
			log.Error(err, " wrong password")
			return domain.UserCore{}, errors.New("wrong password")
		}
		return res, nil
	}
}

// ShowAll implements domain.Service
func (us *userService) ShowAll() ([]domain.UserCore, error) {
	res, err := us.qry.GetAll()
	if err == gorm.ErrRecordNotFound {
		log.Error(err.Error())
		return nil, gorm.ErrRecordNotFound
	} else if err != nil {
		log.Error(err.Error())
		return nil, errors.New(config.DATABASE_ERROR)
	}

	if len(res) == 0 {
		log.Info("no data")
		return nil, errors.New(config.DATA_NOTFOUND)
	}

	return res, nil
}

// Update implements domain.Service
func (us *userService) Update(id uint, email string) (domain.UserCore, error) {
	panic("unimplemented")
}
