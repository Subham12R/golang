package services

import (
	"errors"
	"go-auth/internal/auth"
	"go-auth/internal/dto"
	"go-auth/internal/models"
	"go-auth/internal/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)


const refreshTokenTTL = 7 * 24 * time.Hour

type AuthService struct {
	repo         *repositories.UserRepository
	refreshRepo  *repositories.RefreshTokenRepository
	jwtSecret    string
}

func NewAuthService(
	repo *repositories.UserRepository,
	refreshRepo *repositories.RefreshTokenRepository,
	jwtSecret string,
) *AuthService {
	return &AuthService{
		repo:        repo,
		refreshRepo: refreshRepo,
		jwtSecret:   jwtSecret,
	}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return nil, errors.New("all fields are required")
	}

	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return s.buildAuthResponse(user)
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return s.buildAuthResponse(user)
}


func (s *AuthService) Refresh(refreshToken string) (*dto.AuthResponse, error) {
	stored, err := s.refreshRepo.FindByToken(refreshToken)
	if err != nil {
		return nil, err
	}

	if stored == nil || stored.Revoked || stored.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("invalid refresh token")
	}

	user, err := s.repo.FindByID(stored.UserID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if err := s.refreshRepo.Revoke(refreshToken); err != nil {
		return nil, err
	}

	return s.buildAuthResponse(user)
}


func (s *AuthService) Logout(refreshToken string) error {
	return s.refreshRepo.Revoke(refreshToken)
}


func (s *AuthService) GetUserByID(id int64) (*dto.UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	return &dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}


func (s *AuthService) buildAuthResponse(user *models.User) (*dto.AuthResponse, error) {
	accessToken, err := auth.GenerateAccessToken(user, s.jwtSecret)
	if err != nil {
		return nil, err
	}

	refreshToken, err := auth.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}

	stored := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(refreshTokenTTL),
	}
	if err := s.refreshRepo.Create(stored); err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    900,
	}, nil
}
