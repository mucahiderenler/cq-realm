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

func (r *VillageRepository) GetAllVillages() ([]*models.Village, error) {
	var villages []*models.Village
	rowVillages, err := r.DB.Query("SELECT id, name, x, y, coalesce(owner_name, ''), owner_id, point, village_type FROM villages")

	if err != nil {
		return nil, err
	}

	for rowVillages.Next() {
		village := &models.Village{}
		err := rowVillages.Scan(&village.ID, &village.Name, &village.X, &village.Y, &village.Owner_name, &village.Owner_id, &village.Point, &village.Type)

		if err != nil {
			return nil, err
		}

		villages = append(villages, village)

	}

	return villages, nil
}
