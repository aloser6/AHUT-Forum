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
func InsertLock() {
	mysql.Mu.Lock()
	Insert()
	mysql.Mu.Unlock()
}
func Updata(undatedUser *models.User) {
	user := undatedUser
	db.Save(user)
}
func Select(id uint) (*models.User, error) {
	finduser := &models.User{}
	var finduserId = finduser.ID

	err := db.First(finduserId, id).Error
	if err != nil {
		return nil, err
	}
	return finduser, nil

}
func SelectLock() {
	mysql.Mu.RLock()
	Select()
	mysql.Mu.Unlock()
}
func Delete(id uint) {
	deleteduser := &models.User{}
	var userId = deleteduser.ID
	userId = id

	db.Delete(userId)
}
