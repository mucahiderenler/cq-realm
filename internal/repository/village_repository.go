package repository

import (
	"context"
	"fmt"
	models "mucahiderenler/conquerors-realm/internal/models"

	"github.com/jmoiron/sqlx"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type VillageRepository struct {
	DB *sqlx.DB
}

func NewVillageRepository(db *sqlx.DB) *VillageRepository {
	return &VillageRepository{DB: db}
}

func (r *VillageRepository) GetByID(ctx context.Context, id string) (*models.Village, error) {
	village, err := models.Villages(Load(Rels(models.VillageRels.Buildings)), Where("id = ?", id)).One(ctx, r.DB)
	if err != nil {
		fmt.Println(err)
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

func (r *VillageRepository) GetAllVillages(ctx context.Context) ([]*models.Village, error) {
	var villages []*models.Village
	villages, err := models.Villages().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return villages, nil
}
