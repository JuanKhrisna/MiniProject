package domain

type Service interface {
	InsertData(domain Obat) (response Obat, err error)
	UpdateData(id int) (response Obat, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain Obat)
	Delete(id int) (err error)
	GetByID(id int) (domain Obat, err error)
	GetByName(name string) (domain Obat, err error)
}
