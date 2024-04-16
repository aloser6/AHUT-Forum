package dao

import (
	"AHUT-Forum/config"
	"AHUT-Forum/models"
)

var mysql = config.Mysql{}
var db = mysql.Db

func Insert(newUser *models.User) {
	user := newUser
	db.Create(user)
}
func InsertLock(newUser *models.User) {
	mysql.Mu.Lock()
	Insert(newUser)
	mysql.Mu.Unlock()
}
func Updata(updatedUser *models.User) {
	user := updatedUser
	db.Save(user)
}
func Select(id uint, findUser *models.User) (*models.User, error) {
	findUserId := findUser.ID

	err := db.First(findUserId, id).Error
	if err != nil {
		return nil, err
	}
	return findUser, nil

}
func SelectLock(id uint, findUser *models.User) {
	mysql.Mu.RLock()
	Select(id, findUser)
	mysql.Mu.Unlock()
}
func Delete(id uint, deletedUser *models.User) {
	userId := deletedUser.ID
	userId = id

	db.Delete(userId)
}
