package repository

import (
	"fmt"
	"melodyhub/pkg/models"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}


func (r *Repository) GetSongs(filter string, page int, limit int) ([]models.Song, error) {
	var songs []models.Song
	offset := (page - 1) * limit

	query := r.db.Model(&models.Song{})
	
	if filter != "" {
		query = query.Where("group LIKE ? OR song LIKE ?", "%"+filter+"%", "%"+filter+"%")
	}

	err := query.Offset(offset).Limit(limit).Find(&songs).Error
	return songs, err
}

func (r *Repository) GetSong(id string) (*models.Song, error) {
	var song models.Song

	err := r.db.First(&song, id).Error
	if err != nil {
		return nil, err
	}

	return &song, nil
}

func (r *Repository) CreateSong(song *models.Song) error {
	err := r.db.Create(song).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateSong(id string, song *models.Song) error {
	result := r.db.Model(&models.Song{}).Where("id = ?", id).Updates(&song)

	if result.RowsAffected == 0 {
		return fmt.Errorf("no song found with ID %s", id)
	}

	err := result.Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteSong(id string) error {
	err := r.db.Delete(&models.Song{}, id).Error
	if err != nil {
		return err
	}

	return nil
}