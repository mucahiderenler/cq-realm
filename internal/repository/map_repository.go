package repository

import (
	"database/sql"
	"encoding/json"
	models "mucahiderenler/conquerors-realm/internal/models"
)

type MapRepository struct {
	DB *sql.DB
}

func NewMapRepository(db *sql.DB) *MapRepository {
	return &MapRepository{DB: db}
}

func (repo *MapRepository) GetMap() (*models.Map, error) {
	m := &models.Map{}
	var villagesJSON []byte
	err := repo.DB.QueryRow("select villages from map where id = 1").Scan(&villagesJSON)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(villagesJSON, &m.Villages)

	if err != nil {
		return nil, err
	}
	return m, err
}
