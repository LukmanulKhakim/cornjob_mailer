package domain

type UserCore struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Status   string
	Token    string
	Wallet   WalletCore
	Mailer   MailerCore
}

type WalletCore struct {
	ID_Wallet  uint
	IDCurrency uint
	IDUser     uint
	Amount     float64
}

type CurrencyCore struct {
	ID_Current uint
	Currency   string
}

type MailerCore struct {
	ID_verifikasi uint
	IDUser        uint
	Email         string
	Pin           string
	Status        string
}

type Repository interface {
	Add(data UserCore) (UserCore, error)
	Login(input UserCore) (UserCore, error)
	GetAll() ([]UserCore, error)
	Edit(id uint, email string) (UserCore, error)
}

type Service interface {
	Register(data UserCore) (UserCore, error)
	Login(input UserCore) (UserCore, error)
	ShowAll() ([]UserCore, error)
	Update(id uint, email string) (UserCore, error)
}
