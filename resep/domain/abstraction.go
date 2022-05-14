package domain

type Service interface {
	InsertData(domain Resep) (response Resep, err error)
	UpdateData(id int) (response Resep, err error)
	DeleteData(id int) (err error)
}

type Repository interface {
	Save(domain Resep) (id int, err error)
	Delete(id int) (err error)
	GetByID(id int) (domain Resep, err error)
}
