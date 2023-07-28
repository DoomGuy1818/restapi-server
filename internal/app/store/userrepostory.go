package store

import "github.com/DoomGuy1818/restapi-server/model"

type UserRepository struct {
	store *Store
}

func (s *UserRepository) Create(u *model.User) (*model.User, error) {
	if err := s.store.db.QueryRow("INSERT INTO users(email, encrypted_password) VALUES ($1, $2) RETURNING id",
	u.Email,
	u.EncryptedPassword,
	).Scan(&u.ID); err != nil {
		return nil, err
	}


	return nil, nil
}

func (s *Store) FindByEmail(email string) (*model.User, error){
	return nil, nil
}