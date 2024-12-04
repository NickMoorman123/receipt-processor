package objects

import (
	"gorm.io/gorm"
)

type Receipt struct {
	gorm.Model
	ID					uint   `gorm:"unique;primaryKey;autoIncrement:true"`
	UUID 				string
	Retailer			string `json:"retailer"`
	PurchaseDate		string `json:"purchaseDate"`
	PurchaseTime		string `json:"purchaseTime"`
	Total				string `json:"total"`
	Items				[]Item `json:"items"`
	Points				int
}

type Item struct {
	gorm.Model
	ID					uint   `gorm:"unique;primaryKey;autoIncrement"`
	ShortDescription	string `json:"shortDescription"`
	Price				string `json:"price"`
	ReceiptID			uint
}