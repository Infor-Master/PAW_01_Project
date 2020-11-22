package model

import "github.com/jinzhu/gorm"

type WorkersZone struct { // Lista de workers a trabalhar em cada zona
	gorm.Model `swaggerignore:"true"`
	IdWorker   int "gorm:\"foreignKey:id;references:Worker\""
	IdZone     int "gorm:\"foreignKey:id;references:Zone\""
}
