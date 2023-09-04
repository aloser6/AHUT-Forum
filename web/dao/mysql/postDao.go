package mysql

import (
	"gorm.io/gorm"
)

type PostDB struct {
	db *gorm.DB
}
