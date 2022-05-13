package domain

type Service interface {
	CreateToken(username, password string) (token string, err error)
	InsertData(domain User) (response User, err error)
	UpdateData(id int) (response User, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain User) (id int, err error)
	Delete(id int) (err error)
	GetByUsernamePassword(username, password string) (domain User, err error)
	GetByID(id int) (domain User, err error)
}
