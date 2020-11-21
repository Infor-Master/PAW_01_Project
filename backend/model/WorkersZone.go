package model

import "github.com/jinzhu/gorm"

type WorkersZone struct { // Lista de workers a trabalhar em cada zona
	gorm.Model `swaggerignore:"true"`
	IdWorker   Worker "gorm:\"foreignKey:id\""
	IdZone     Zone "gorm:\"foreignKey:id\""
}
