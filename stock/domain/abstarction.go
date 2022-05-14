package domain

type Service interface {
	InsertData(domain Stock) (response Stock, err error)
	UpdateData(id int) (response Stock, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain Stock) (id int, err error)
	Delete(id int) (err error)
	GetByName(name string) (domain Stock, err error)
	GetByID(id int) (domain Stock, err error)
}
