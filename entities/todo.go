package entities

import (
	"gorm.io/gorm"
)

type Todo struct {
	Owner	   	string	`json:"owner"`
	Title       string	`json:"title"`
	Completed   bool  	`json:"complete" gorm:"default:false"`
	DueDate     string	`json:"duedate"`
	gorm.Model
}