package users

import (
	"Web-chat/pkg/services"
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
	err = services.BcryptHasher{}.CheckPasswordHash(password, user.Password)
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

func (u *UserRepository) Find(name string) ([]User, error) {
	var user []User
	err := u.DB.Where("name LIKE ? LIMIT 20", "%"+name+"%").Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
