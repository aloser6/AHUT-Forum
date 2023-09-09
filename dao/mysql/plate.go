package mysql

type Announcements struct {
	AnnouncementId          int    `gorm:"column:announcement_id"` //公告ID
	AnnouncementName        string `gorm:"column:announcement_name"`
	AnnouncementTime        string `gorm:"column:announcement_time"`
	AnnouncementInformation string `gorm:"column:announcement_information"`
	PlateId                 int    `gorm:"column:plate_id"`
}

type PlateFans struct {
	PlateId    int    `gorm:"column:plate_id"`
	FansNumber int    `gorm:"column:fans_number"`
	FansID     string `gorm:"column:fans_id"`
}

type EasyManage struct {
	PlateId      int    `gorm:"column:plate_id"`
	ManageNumber int    `gorm:"column:manage_number"`
	ManageID     string `gorm:"column:manage_id"` //管理者们的ID
}

type Plates struct {
	PlateId     int    `gorm:"column:plate_id"`
	PlateName   string `gorm:"column:plate_name"`
	Views       int    `gorm:"column:views"`
	PlateType   string `gorm:"column:plate_type"`
	PostsNumber int    `gorm:"column:posts_number"`
	PlateTime   string `gorm:"column:plate_time"`
	ModeratorId int    `gorm:"column:moderator_id"` //版主ID
}

type EasyPosts struct {
	PostId     int    `gorm:"column:post_id"`
	PostName   string `gorm:"column:post_name"`
	Author     string `gorm:"column:author"`
	Views      int    `gorm:"column:views"`
	PostType   string `gorm:"column:post_type"`
	LikeNumber int    `gorm:"column:like_number"`
	PostTime   string `gorm:"column:post_time"`
	PlateId    int    `gorm:"column:plate_id"`
}
