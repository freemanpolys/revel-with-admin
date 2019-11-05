package models

import "github.com/jinzhu/gorm"

type GroupLabel struct {
	gorm.Model
	Name string
}
