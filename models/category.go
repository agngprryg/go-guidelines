package models

import "time"

type Category struct {
	ID			uint			`gorm:"primaryKey"`
	Name     	string    		`gorm:"type:varchar(100);not null" json:"name"`
	Products 	[]Product 		`gorm:"foreignKey:CategoryID" json:"products,omitempty"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}