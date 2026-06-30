package repositories

import (
	"database/sql"
	"errors"
	"go-auth/internal/models"
)

type RefreshTokenRepository struct {
	db *sql.DB
}

func NewRefreshTokenRepository(db *sql.DB) *RefreshTokenRepository {
	return &RefreshTokenRepository{
		db: db,
	}
}

func (r *RefreshTokenRepository) Create(token *models.RefreshToken) error {
	query := `
		INSERT INTO refresh_tokens (
			user_id,
			token,
			expires_at
		)
		VALUES ($1, $2, $3)
		RETURNING id, revoked, created_at;
	`

	return r.db.QueryRow(
		query,
		token.UserID,
		token.Token,
		token.ExpiresAt,
	).Scan(
		&token.ID,
		&token.Revoked,
		&token.CreatedAt,
	)
}


func (r *RefreshTokenRepository) FindByToken(token string) (*models.RefreshToken, error) {
	query := `
		SELECT
			id,
			user_id,
			token,
			expires_at,
			revoked,
			created_at
		FROM refresh_tokens
		WHERE token = $1;
	`

	var rt models.RefreshToken

	err := r.db.QueryRow(query, token).Scan(
		&rt.ID,
		&rt.UserID,
		&rt.Token,
		&rt.ExpiresAt,
		&rt.Revoked,
		&rt.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &rt, nil
}

func (r *RefreshTokenRepository) Revoke(token string) error {
	query := `
		UPDATE refresh_tokens
		SET revoked = TRUE
		WHERE token = $1;
	`

	_, err := r.db.Exec(query, token)
	return err
}
