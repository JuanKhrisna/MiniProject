package domain

type Service interface {
	InsertData(domain Apotik) (response Apotik, err error)
	UpdateData(id int) (response Apotik, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain Apotik) (id int, err error)
	Delete(id int) (err error)
	GetByName(name string) (domain Apotik, err error)
	GetByID(id int) (domain Apotik, err error)
	GetByAlamat(alamat string) (domain Apotik, err error)
}
