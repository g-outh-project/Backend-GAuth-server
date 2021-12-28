package method

import (
	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/model"
)

func InsertUser(user dto.SignupReq) error {
	db := db.GetDB()
	tx := db.Create(&model.User{
		Id:       user.Id,
		Password: user.Password,
		Name:     user.Name,
		Email:    user.Email,
		School:   user.School,
		Birth:    user.Birth,
		Nickname: user.Nickname,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectUserById(id string) (model.User, error) {
	var user model.User
	db := db.GetDB()
	tx := db.Find(&user, model.User{Id: id})
	if tx.Error != nil {
		return user, tx.Error
	}
	return user, nil
}

func SelectUser() ([]model.User, error) {
	var users []model.User
	db := db.GetDB()
	tx := db.Select("*").Find(&users)
	if tx.Error != nil {
		return users, tx.Error
	}
	return users, nil
}
