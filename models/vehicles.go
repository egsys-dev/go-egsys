package models

type Vehicle struct {
	Vehicle_id uint    `gorm:"primaryKey;" json:"id"`
	Fleet_id   int     `gorm:"index"`
	Name       string  `gorm:"not null" json:"name"`
	Max_speed  float32 `gorm:"type:decimal(6,2)" json:"max_speed"`
	//FleetAlert []FleetAlert `gorm:"foreignKey:Fleet_id"`
}

var Vehicles []Vehicle
