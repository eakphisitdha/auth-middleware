package repository

import (
	"app/model"
	"database/sql"
)

type IRepository interface {
	Get(req model.GetRequest) ([]model.Data, error)
	Add(req model.AddRequest) error
	Update(req model.UpdateRequest) error
	Delete(req model.DeleteRequest) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) IRepository {
	return &Repository{db: db}
}

func (r *Repository) Get(req model.GetRequest) ([]model.Data, error) {
	//
	//SQL logic
	//
	data := []model.Data{}
	return data, nil
}

func (r *Repository) Add(req model.AddRequest) error {
	//
	//SQL logic
	//

	return nil
}

func (r *Repository) Update(req model.UpdateRequest) error {
	//
	//SQL logic
	//

	return nil
}

func (r *Repository) Delete(req model.DeleteRequest) error {
	//
	//SQL logic
	//

	return nil
}
