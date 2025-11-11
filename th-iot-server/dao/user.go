package dao

import (
	"th-iot-server/models"
	"th-iot-server/utils"
)

type UserDAO struct{}

var User = &UserDAO{}

func (d *UserDAO) Create(user *models.User) error {
	return utils.DB.Create(user).Error
}

func (d *UserDAO) Update(user *models.User) error {
	return utils.DB.Save(user).Error
}

func (d *UserDAO) Delete(user *models.User) error {
	return utils.DB.Delete(user).Error
}

func (d *UserDAO) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := utils.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) FindByUsername(username string) (*models.User, error) {
	var user models.User
	if err := utils.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := utils.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (d *UserDAO) FindAll() ([]models.User, error) {
	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDAO) FindByCondition(cond map[string]interface{}) ([]models.User, error) {
	var users []models.User
	if err := utils.DB.Where(cond).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
