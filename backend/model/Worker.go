package model

import "github.com/jinzhu/gorm"

type Worker struct {
	gorm.Model `swaggerignore:"true"`
	Username   string "gorm:unique"
	Password   string
	Name       string 
	Admin      bool   `gorm:"default=false"`
	Zones []*Zone `gorm:"many2many:worker_zone;"`
	//ZoneID uint
}
