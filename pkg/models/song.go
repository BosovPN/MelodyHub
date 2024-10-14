package models

type Song struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Group       string `json:"group" gorm:"not null"`
	Song        string `json:"song" gorm:"not null"`
	ReleaseDate string `json:"releaseDate" gorm:"not null"`
	Text        string `json:"text" gorm:"not null"`
	Link        string `json:"link" gorm:"not null"`
}
