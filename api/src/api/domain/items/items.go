package items

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	Thumbnail   string  `json:"thumbnail"`
}
