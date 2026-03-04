package models

type Product struct {
	ID 			uint		`gorm:"primaryKey"`	
	Name        string   	`gorm:"type:varchar(150);not null" json:"name"`
	Description string   	`gorm:"type:text" json:"description"`
	Price       float64  	`gorm:"type:decimal(15,2);not null" json:"price"`
	Stock       int      	`gorm:"default:0" json:"stock"`
	CategoryID  uint     	`gorm:"not null" json:"category_id"`
	Category    Category 	`gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}