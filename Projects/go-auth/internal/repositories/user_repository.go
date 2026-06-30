package repositories

import (
	"database/sql"
	"errors"
	"go-auth/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository{
	return &UserRepository{
		db: db,
	}
}

func(r *UserRepository) Create(user *models.User) error {
	query := `
		INSERT INTO users (
			name,
			email,
			password_hash
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at;	
	`

	return r.db.QueryRow(
		query,
		user.Name,
		user.Email,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
}


func (r *UserRepository) FindByEmail(email string) (*models.User, error){
	query := `
		SELECT
			id, 
			name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE email = $1;
	`
	var user models.User

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,		
		&user.Name,		
		&user.Email,		
		&user.PasswordHash,		
		&user.CreatedAt,		
		&user.UpdatedAt,	
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	query := `
		SELECT
			id,
			name,
			email,
			password_hash,
			created_at,
			updated_at
		FROM users
		WHERE id = $1;
	`

	var user models.User

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}