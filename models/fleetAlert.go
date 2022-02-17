package models

type FleetAlert struct {
	Fleet_alert_id uint `gorm:"primaryKey;" json:"id"`
	Fleet_id       int  `gorm:"index"`
	// Fleet          []Fleet `gorm:"foreignKey:Fleet_id"`
	Webhook string `gorm:"not null" json:"webhook"`
}

var FleetAlerts []FleetAlert
