package method

import (
	"github.com/Backend-GAuth-server/db"
	"github.com/Backend-GAuth-server/model"
)

func InsertClient(clientId string, secret string) error {

	db := db.GetDB()
	tx := db.Create(&model.Client{
		ClientId: clientId,
		Secret:   secret,
	})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func SelectKeyByCid(cid string) (model.Client, error) {
	var key model.Client
	db := db.GetDB()
	tx := db.Find(&key, model.Client{ClientId: cid})
	if tx.Error != nil {
		return key, tx.Error
	}
	return key, nil
}
