package mysql

/*一个板块的管理员*/
type Manage struct {
	PlateId   int    `gorm:"column:plate_id"`
	AccountId string `gorm:"column:account_id"`
}

/*一个人是多少个板块的版主*/
type Moderators struct {
	AccountId int    `gorm:"column:account_id"`
	PlatesId  string `gorm:"column:plates_id"`
}

/*系统管理员*/
type SystemAdministrators struct {
	AccountId int `gorm:"column:account_id"`
}
