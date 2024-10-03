package repository

import (
	"context"
	"mucahiderenler/conquerors-realm/internal/models"

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
		return nil, err
	}
	return village, nil
}

func (r *VillageRepository) GetAllVillages(ctx context.Context) ([]*models.Village, error) {
	var villages []*models.Village
	villages, err := models.Villages().All(ctx, r.DB)
	if err != nil {
		return nil, err
	}

	return villages, nil
}
