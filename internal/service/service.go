package service

import (
	"melodyhub/internal/repository"
	"melodyhub/pkg/models"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repo: repo}
}


func (s *Service) GetSongs(filter string, page int, limit int) ([]models.Song, error) {
	songs, err := s.repo.GetSongs(filter, page, limit)
	if err != nil {
		return nil, err
	}

	return songs, nil
}

func (s *Service) GetSong(id string) (*models.Song, error) {
	song, err := s.repo.GetSong(id)
	if err != nil {
		return nil, err
	}

	return song, nil
}

func (s *Service) AddSong(song *models.Song) error {
	err := s.repo.CreateSong(song)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateSong(id string, song *models.Song) error {
	err := s.repo.UpdateSong(id, song)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSong(id string) error {
	err := s.repo.DeleteSong(id)
	if err != nil {
		return err
	}

	return nil
}