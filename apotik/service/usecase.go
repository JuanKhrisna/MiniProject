package service

import (
	"miniproject/apotik/domain"
	errorConv "miniproject/helper/error"
)

type apotikService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (us apotikService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)
	return errorConv.Conversion(errResp)
}

// InsertData implements domain.Service
func (apotikService) InsertData(domain domain.Apotik) (response domain.Apotik, err error) {
	panic("unimplemented")
}

// UpdateData implements domain.Service
func (apotikService) UpdateData(id int) (response domain.Apotik, err error) {
	panic("unimplemented")
}

func NewApotikService(repo domain.Repository) domain.Service {
	return apotikService{
		repository: repo,
	}
}
