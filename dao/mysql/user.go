package mysql

type Accounts struct {
	AccountId int    `gorm:"column:account_id"`
	Account   string `gorm:"column:account"`
	Password  string `gorm:"column:password"`
	Email     string `gorm:"column:email"`
}

type ConcernPeople struct {
	AccountId           int    `gorm:"column:account_id"`
	ConcernPeopleNumber int    `gorm:"column:concern_people_number"`
	ConcernPeopleID     string `gorm:"column:concern_people_id"`
}

type ConcernPlates struct {
	AccountId           int    `gorm:"column:account_id"`
	ConcernPlatesNumber int    `gorm:"column:concern_plates_number"`
	ConcernPlatesID     string `gorm:"column:concern_plates_id"`
}

type UserFans struct {
	AccountId  int    `gorm:"column:account_id"`
	FansNumber int    `gorm:"column:fans_number"`
	FansID     string `gorm:"column:fans_id"`
}

type Information struct {
	AccountId   int    `gorm:"column:account_id"`
	AccountName string `gorm:"column:account_name"`
	Age         int    `gorm:"column:age"`
	Sex         string `gorm:"column:sex"`
	PostsNumber int    `gorm:"column:posts_number"`
	PostsId     string `gorm:"column:posts_id"`
}
