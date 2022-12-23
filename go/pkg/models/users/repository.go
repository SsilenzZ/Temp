package users

import (
	"Web-chat/pkg/service"
	"gorm.io/gorm"
)

type UserRepositoryI interface{}

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (u UserRepository) GetHashedPassword(email, password string) (int, error) {
	var user User

	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0, err
	}
	err = service.BcryptHasher{}.CheckPasswordHash(password, user.Password)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u UserRepository) CreateUser(email, password string, name string) bool {
	var user User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err == nil {
		if user.Password != "" {
			return false
		}
		user.Password = password
		name = "test"
		user.Name = name
		u.DB.Save(&user)
		return true
	}
	user.Email = email
	user.Password = password
	user.Name = name
	u.DB.Select("email", "password", "name").Create(&user)
	return true
}
