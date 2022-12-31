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

func (u UserRepository) Find(name string, id int) ([]User, error) {
	var user []User
	err := u.DB.Where("name LIKE ? AND id != ? LIMIT 20", "%"+name+"%", id).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepository) GetById(id int) (int, string, error) {
	var user User
	err := u.DB.Where("id = ?", id).Take(&user).Error
	return user.ID, user.Name, err
}
