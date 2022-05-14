package service

import (
	errorConv "miniproject/helper/error"
	"miniproject/resep/domain"
)

type resepService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (us resepService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)
	return errorConv.Conversion(errResp)
}

// InsertData implements domain.Service
func (resepService) InsertData(domain domain.Resep) (response domain.Resep, err error) {
	panic("unimplemented")
}

// UpdateData implements domain.Service
func (resepService) UpdateData(id int) (response domain.Resep, err error) {
	panic("unimplemented")
}

func NewResepService(repo domain.Repository) domain.Service {
	return resepService{
		repository: repo,
	}
}
