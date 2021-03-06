package service

import (
	errorConv "miniproject/helper/error"
	"miniproject/user/domain"
)

type userService struct {
	repository domain.Repository
}

// CreateToken implements domain.Service
func (userService) CreateToken(username string, password string) (token string, err error) {
	panic("unimplemented")
}

// DeleteData implements domain.Service
func (us userService) DeleteData(id int) (err error) {
	errResp := us.repository.Delete(id)
	return errorConv.Conversion(errResp)
}

// InsertData implements domain.Service
func (us userService) InsertData(domain domain.User) (response domain.User, err error) {
	panic("unimplemented")
}

// UpdateData implements domain.Service
func (userService) UpdateData(id int) (response domain.User, err error) {
	panic("unimplemented")
}

func NewUserService(repo domain.Repository) domain.Service {
	return userService{
		repository: repo,
	}
}
