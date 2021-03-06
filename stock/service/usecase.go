package service

import (
	errorConv "miniproject/helper/error"
	"miniproject/stock/domain"
)

type stockService struct {
	repository domain.Repository
}

// DeleteData implements domain.Service
func (us stockService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)
	return errorConv.Conversion(errResp)
}

// InsertData implements domain.Service
func (stockService) InsertData(domain domain.Stock) (response domain.Stock, err error) {
	panic("unimplemented")
}

// UpdateData implements domain.Service
func (stockService) UpdateData(id int) (response domain.Stock, err error) {
	panic("unimplemented")
}

func NewStockService(repo domain.Repository) domain.Service {
	return stockService{
		repository: repo,
	}
}
