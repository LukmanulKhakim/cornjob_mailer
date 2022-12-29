package delivery

import "cornjobmailer/features/user/domain"

type RegisterFormat struct {
	Name     string `json:"name" form:"name" validate:"required,min=4,max=30"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password"`
}

type LoginFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func ToDomain(i interface{}) domain.UserCore {
	switch i.(type) {
	case RegisterFormat:
		cnv := i.(RegisterFormat)
		return domain.UserCore{Name: cnv.Name, Email: cnv.Email, Password: cnv.Password}
	case LoginFormat:
		cnv := i.(LoginFormat)
		return domain.UserCore{Email: cnv.Email, Password: cnv.Password}
	}

	return domain.UserCore{}
}
