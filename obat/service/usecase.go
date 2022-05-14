package service

import "miniproject/obat/domain"

type obatService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (obatService) DeleteData(id int) (err error) {
	panic("unimplemented")
}

// InsertData implements domain.Service
func (obatService) InsertData(domain domain.Obat) (response domain.Obat, err error) {
	panic("unimplemented")
}

// UpdateData implements domain.Service
func (obatService) UpdateData(id int) (response domain.Obat, err error) {
	panic("unimplemented")
}

func NewObatService(repo domain.Repository) domain.Service {
	return obatService{
		repository: repo,
	}
}
