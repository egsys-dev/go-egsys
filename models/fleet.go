package models

type Fleet struct {
	Fleet_id   uint         `gorm:"primaryKey;" json:"id"`
	Name       string       `gorm:"not null" json:"name"`
	Max_speed  float32      `gorm:"type:decimal(6,2)" json:"max_speed"`
	FleetAlert []FleetAlert `gorm:"foreignKey:Fleet_id"`
	Vehicle    []Vehicle    `gorm:"foreignKey:Fleet_id"`
}

var Fleets []Fleet
