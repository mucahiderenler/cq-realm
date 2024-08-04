package services

import (
	"errors"
	"mucahiderenler/conquerors-realm/internal/models"
	"mucahiderenler/conquerors-realm/internal/repository"
	"regexp"

	"go.uber.org/zap"
)

type VillageService struct {
	Repo   *repository.VillageRepository
	Logger *zap.Logger
}

func NewVillageService(repo *repository.VillageRepository, logger *zap.Logger) *VillageService {
	return &VillageService{Repo: repo, Logger: logger}
}

// Validate village data
func (s *VillageService) validateVillage(village *models.Village) error {
	if village.Name == "" {
		return errors.New("village name is required")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9 ]+$`).MatchString(village.Name) {
		return errors.New("village name can only contain letters, numbers, and spaces")
	}
	// Additional validations can be added here
	return nil
}

func (s *VillageService) GetAllVillages() ([]*models.Village, error) {
	villages, err := s.Repo.GetAllVillages()

	if err != nil {
		return nil, err
	}

	return villages, nil
}

// Get a village by its ID
func (s *VillageService) GetVillageByID(id string) (*models.Village, error) {
	village, err := s.Repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	// Perform any additional business logic if needed
	return village, nil
}

// Create a new village
func (s *VillageService) CreateVillage(village *models.Village) error {
	// Validate the village data
	if err := s.validateVillage(village); err != nil {
		return err
	}

	// Set initial resources or other properties if needed
	// village.Resources = models.Resources{
	// 	Wood:  100,
	// 	Stone: 100,
	// 	Food:  100,
	// }

	return s.Repo.Create(village)
}

// Update an existing village
func (s *VillageService) UpdateVillage(village *models.Village) error {
	// Validate the village data
	if err := s.validateVillage(village); err != nil {
		return err
	}

	// Additional business logic or calculations can be added here

	return s.Repo.Update(village)
}

// Delete a village by its ID
func (s *VillageService) DeleteVillage(id string) error {
	// Perform any cleanup or additional checks before deletion
	return s.Repo.Delete(id)
}
