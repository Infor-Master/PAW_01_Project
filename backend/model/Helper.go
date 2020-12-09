package model

type Helper struct {
	Name       string "gorm:unique"
	Latitude   float32 "gorm:not null"
	Longitude  float32 "gorm:not null"
	Limits      int "gorm:not null"
	PplCount   int `gorm:"default=0"`
	Workers []*Worker `gorm:"many2many:worker_zone;"`
}
