package repository

import (
	"cornjobmailer/features/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Password string
	Status   string
	Token    string `gorm:"-:migration;<-:false"`
	Wallet   Wallet `gorm:"foreignKey:IDUser"`
	Mailer   Mailer `gorm:"foreignKey:IDUser"`
}

type Wallet struct {
	gorm.Model
	IDCurrency uint
	IDUser     uint
	Amount     float64
}

type Currency struct {
	gorm.Model
	Currency string
	Wallet   Wallet `gorm:"foreignKey:IDCurrency"`
}

type Mailer struct {
	gorm.Model
	IDUser uint
	Email  string
	Pin    string
	Status string
}

var (
	uw domain.WalletCore
	um domain.MailerCore
)

func FromDomain(uc domain.UserCore) User {
	return User{
		Model:    gorm.Model{ID: uc.ID},
		Name:     uc.Name,
		Email:    uc.Email,
		Password: uc.Password,
		Status:   uc.Status,
		Token:    uc.Token,
		Wallet: Wallet{
			Model:      gorm.Model{ID: uw.ID_Wallet},
			IDCurrency: uc.Wallet.IDCurrency,
			IDUser:     uc.Wallet.IDUser,
			Amount:     uc.Wallet.Amount},
		Mailer: Mailer{
			Model:  gorm.Model{ID: um.ID_verifikasi},
			IDUser: uc.Mailer.IDUser,
			Email:  uc.Mailer.Email,
			Pin:    uc.Mailer.Pin,
			Status: uc.Mailer.Status},
	}
}

func ToDomain(u User) domain.UserCore {
	return domain.UserCore{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Status:   u.Status,
		Token:    u.Token,
		Wallet: domain.WalletCore{
			IDCurrency: u.Wallet.IDCurrency,
			IDUser:     u.Wallet.IDUser,
			Amount:     u.Wallet.Amount},
		Mailer: domain.MailerCore{
			IDUser: um.IDUser,
			Email:  u.Email,
			Pin:    um.Pin,
			Status: um.Status},
	}
}

func ToDomainArray(au []User) []domain.UserCore {
	var res []domain.UserCore
	for _, val := range au {
		res = append(res, domain.UserCore{
			ID:       val.ID,
			Name:     val.Name,
			Email:    val.Email,
			Password: val.Password,
			Status:   val.Status,
			Token:    val.Token,
			Wallet: domain.WalletCore{
				IDCurrency: val.Wallet.IDCurrency,
				IDUser:     val.Wallet.IDUser,
				Amount:     val.Wallet.Amount},
			Mailer: domain.MailerCore{
				IDUser: um.IDUser,
				Email:  val.Email,
				Pin:    um.Pin,
				Status: um.Status},
		})
	}
	return res
}
