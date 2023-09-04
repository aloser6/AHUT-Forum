package mysql

type Comments struct {
	CommentId   int    `gorm:"column:comment_id"` //评论ID
	Comment     string `gorm:"column:comment"`
	LikeNumber  int    `gorm:"column:like_number"` //点赞数
	CommentTime string `gorm:"column:comment_time"`
	PostId      int    `gorm:"column:post_id"`
	AccountId   int    `gorm:"column:account_id"`
}

type Posts struct {
	PostId         int    `gorm:"column:post_id"`
	PostName       string `gorm:"column:post_name"`
	Views          int    `gorm:"column:views"`
	PostType       string `gorm:"column:post_type"`
	LikeNumber     int    `gorm:"column:like_number"`
	CommentsNumber int    `gorm:"column:comment_number"` //评论数
	ForwardNumber  int    `gorm:"column:forward_number"` //转发数
	PostTime       string `gorm:"column:post_time"`
	AccountId      int    `gorm:"column:account_id"`
}

type PostDetails struct {
	PostId  int    `gorm:"column:post_id"`
	Details string `gorm:"column:details"`
}
