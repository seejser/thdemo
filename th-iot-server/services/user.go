package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"th-iot-server/dao"
	"th-iot-server/models"
)

func RegisterUser(username, password, email string) error {
	if _, err := dao.User.FindByUsername(username); err == nil {
		return errors.New("用户已存在")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := models.User{
		Username: username,
		Password: string(hash),
		Email:    email,
	}
	return dao.User.Create(&user)
}

func LoginUser(username, password string) (*models.User, error) {
	user, err := dao.User.FindByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}
	return user, nil
}

// GetUserByID 获取用户信息
func GetUserByID(id uint) (*models.User, error) {
	return dao.User.FindByID(id)
}