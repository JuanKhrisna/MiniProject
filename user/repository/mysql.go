package repository

import (
	"miniproject/user/domain"

	"gorm.io/gorm"
)

type userRepo struct {
	DB gorm.DB
}

// Delete implements domain.Repository
func (ur userRepo) Delete(id int) (err error) {
	return ur.DB.Delete("id = ?", id).Error
}

// GetByID implements domain.Repository
func (ur userRepo) GetByID(id int) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Find("id = ?", id).First(&newRecord).Error
	return toDomain(newRecord), err
}

// GetByUsernamePassword implements domain.Repository
func (ur userRepo) GetByUsernamePassword(username string, password string) (domain domain.User, err error) {
	var newRecord User
	err = ur.DB.Find("username = ?", username).First(&newRecord).Error
	return toDomain(newRecord), err
}

// Save implements domain.Repository
func (ur userRepo) Save(domain domain.User) (id int, err error) {
	record := fromDomain(domain)
	err = ur.DB.Create(&record).Error
	return int(record.ID), err
}

// Update implements domain.Repository
func (ur userRepo) Update(id int, domain domain.User) (err error) {
	record := fromDomain(domain)
	err = ur.DB.Where("id = ?", id).Updates(&record).Error
	return err
}

func NewUserRepository(db gorm.DB) domain.Repository {
	return userRepo{
		DB: db,
	}
}
