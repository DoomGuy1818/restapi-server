package store

import "github.com/DoomGuy1818/restapi-server/model"

type UserRepository struct {
	store *Store
}

func (s *UserRepository) Create(u *model.User) (*model.User, error) {
	_, err := s.store.db.Exec("INSERT INTO users(email, encrypted_password) VALUES (?, ?)",
	u.Email,
	u.EncryptedPassword,
	) 
	if err != nil {
		return nil, err
	}
	if err := s.store.db.QueryRow("SELECT LAST_INSERT_ID(?)", u.ID).Scan(&u.ID);
	err != nil{
		return nil, err
	}
	

	return u, nil
}

func (s *UserRepository) FindByEmail(email string) (*model.User, error){
	u := &model.User{
	}
	if err := s.store.db.QueryRow(
		"SELECT id, email, encrypted_password FROM users WHERE email = ?", 
		u.Email,
	).Scan(
		&u.ID, 
		&u.Email, 
		&u.EncryptedPassword,
		); err !=nil{
			return nil, err
		}
	return u, nil
}