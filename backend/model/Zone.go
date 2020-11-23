package model

import "github.com/jinzhu/gorm"

type Zone struct {
	gorm.Model `swaggerignore:"true"`
	Name       string "gorm:unique"
	Latitude   float32
	Longitude  float32
	Limits      int "gorm:not null"
	PplCount   int `gorm:"default=0"`
	Workers []*Worker `gorm:"many2many:worker_zone;"`
	//Workers []Worker `gorm:"foreignKey:ZoneID"`
}
