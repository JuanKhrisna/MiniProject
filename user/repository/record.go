package repository

import (
	"miniproject/user/domain"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int
	Username string
	Password string
	Email    string
	Alamat   string
}

func toDomain(rec User) domain.User {
	return domain.User{
		Id:       rec.Id,
		Username: rec.Username,
		Password: rec.Password,
		Email:    rec.Email,
		Alamat:   rec.Alamat,
	}
}

func fromDomain(rec domain.User) User {
	return User{
		Id:       rec.Id,
		Username: rec.Username,
		Password: rec.Password,
		Email:    rec.Email,
		Alamat:   rec.Alamat,
	}
}
