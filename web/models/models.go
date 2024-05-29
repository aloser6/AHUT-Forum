package models

import "time"

type User struct {
	UID       int       `gorm:"primaryKey;autoIncrement;column:uid"`
	Account   int       `gorm:"not null;column:account"`
	Password  string    `gorm:"not null;column:password"`
	Username  string    `gorm:"not null;column:username"`
	StartTime time.Time `gorm:"not null;column:starttime"`
	Sex       SexType   `gorm:"not null;column:sex"`
	Grade     string    `gorm:"not null;column:grade"`
	College   string    `gorm:"size:128;column:college"`
	Major     string    `gorm:"size:128;column:major"`
}

type Forum struct {
	FID         int    `gorm:"primaryKey;autoIncrement;column:fid"`
	UID         int    `gorm:"references:users;foreignKey:UID;column:uid"`
	ForumName   string `gorm:"not null;column:forumname"`
	MemberCount int    `gorm:"default:1;column:membercount"`
	Description string `gorm:"type:text;column:description"`
}

type Post struct {
	PID         int       `gorm:"primaryKey;autoIncrement;column:pid"`
	FID         int       `gorm:"references:forums;foreignKey:FID;column:fid"`
	UID         int       `gorm:"references:users;foreignKey:UID;column:uid"`
	Title       string    `gorm:"not null;column:title"`
	Text        string    `gorm:"type:text;not null;column:text"`
	Browse      int       `gorm:"default:0;column:browse"`
	Comment     int       `gorm:"default:0;column:comment"`
	Votes       int       `gorm:"default:0;column:votes"`
	ReleaseTime time.Time `gorm:"not null;column:releasetime"`
}

type Comment struct {
	CID         int       `gorm:"primaryKey;autoIncrement;column:cid"`
	UID         int       `gorm:"references:users;foreignKey:UID;column:uid"`
	PID         int       `gorm:"references:posts;foreignKey:PID;column:pid"`
	FatherCID   int       `gorm:"default:-1;column:fathercid"`
	Votes       int       `gorm:"default:0;column:votes"`
	Text        string    `gorm:"type:text;not null;column:text"`
	ReleaseTime time.Time `gorm:"not null;column:releasetime"`
}

type Follow struct {
	FollowID  int `gorm:"primaryKey;column:followid"`
	UID       int `gorm:"references:users;foreignKey:UID;column:uid"`
	Following int `gorm:"not null;column:following"`
}

type SexType string

const (
	Male   SexType = "男"
	Female SexType = "女"
	Secret SexType = "保密"
)
