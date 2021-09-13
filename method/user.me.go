package method

import (
	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/dto"
	"github.com/Backend-GAuth-server/model"
	"github.com/Backend-GAuth-server/utils"
)

func InsertUser(user dto.SignupReq) {
	db := db.GetDB()
	db.Create(&model.User{
		Id:       user.Id,
		Password: utils.Hash(user.Password),
		Name:     user.Name,
		Email:    user.Email,
		School:   user.School,
		Birth:    user.Birth,
		Nickname: user.Nickname,
	})
}

func SelectUser() []model.User {
	var users []model.User
	db := db.GetDB()
	db.Select("*").Find(&users)
	return users
}
