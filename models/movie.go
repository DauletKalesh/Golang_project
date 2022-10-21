package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Name        string `json:"name"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
	//Director *Director `json:"director"`
}
