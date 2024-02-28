package main

import "database/sql"

type Store interface {
	//Users
	CreateUser() error
}

type Storage struct {
	db *sql.DB
}

func NewStore(db *sqlDB) *Storage {
	return &Storage{
		db,
	}
}

func (s *Storage) CreateUser() {
	return nil
}
