package repository

import (
	"database/sql"
	"mucahiderenler/conquerors-realm/internal/models"
)

type VillageRepository struct {
	DB *sql.DB
}

func NewVillageRepository(db *sql.DB) *VillageRepository {
	return &VillageRepository{DB: db}
}

func (r *VillageRepository) GetByID(id string) (*models.Village, error) {
	village := &models.Village{}
	err := r.DB.QueryRow("SELECT id, name, x, y FROM villages WHERE id = $1", id).Scan(&village.ID, &village.Name, &village.X, &village.Y)
	if err != nil {
		return nil, err
	}
	return village, nil
}

func (r *VillageRepository) Create(village *models.Village) error {
	_, err := r.DB.Exec("INSERT INTO villages (id, name) VALUES ($1, $2)", village.ID, village.Name)
	return err
}

func (r *VillageRepository) Update(village *models.Village) error {
	_, err := r.DB.Exec("UPDATE villages SET name = $1 WHERE id = $2", village.Name, village.ID)
	return err
}

func (r *VillageRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM villages WHERE id = $1", id)
	return err
}
