package store

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
)

type Store struct {
	config *Config
	db     *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: *&config,
	}
}

func (s *Store) Open() error {
	db, err := sql.Open("mysql", s.config.DatabaseURL)
	if err != nil{
		return err
	}

	if err := db.Ping(); err !=nil{
		return err
 	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}


func (s *Store) User() *UserRepository {
	if s.userRepository !=nil {
		return s.userRepository
	}
	
	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}