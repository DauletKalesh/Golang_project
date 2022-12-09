package models

type Movie struct {
	ID          uint        `gorm:"primaryKey"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Genre       []Genre     `json:"genre" gorm:"many2many:movie_genres;"`
	Actor       []Celebrity `json:"actor" gorm:"many2many:movie_celebrities;"`

	Country    int     `json:"country"`
	Rating     float32 `json:"rating"`
	ViewAmount int     `json:"view_amount"`
	Bookmark   []Bookmark
	// ReleaseDate date    `json:"release_date"`
	//Director *Director `json:"director"`
}

type Genre struct {
	ID    uint    `gorm:"primaryKey"`
	Name  string  `json:"name"`
	Movie []Movie `json:"genre" gorm:"many2many:movie_genres;"`
}

type Bookmark struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint
	MovieID     uint
	is_favorite bool
	is_watched  bool
}
