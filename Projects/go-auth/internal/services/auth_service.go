package services

import (
	"errors"
	"go-auth/internal/auth"
	"go-auth/internal/dto"
	"go-auth/internal/models"
	"go-auth/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.UserRepository
	jwtSecret string
}

func newAuthService(repo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		repo: repo,
		jwtSecret: jwtSecret,
	}
}

func(s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error){
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
		Name:  req.Name,
		Email: req.Email,
		PasswordHash: string(hashedPassword),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	token, err := auth.GenerateAccessToken(
		user,
		s.jwtSecret,
	)

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
		AccessToken: token,
		ExpiresIn: 900,
	}, nil
}

func(s *AuthService) Login(req * dto.LoginRequest) (*dto.AuthResponse, error){
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil{
		return nil, err
	}

	if user == nil{
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid Credentials")
	}

	token, err := auth.GenerateAccessToken(
		user,
		s.jwtSecret,
	)

	return &dto.AuthResponse{
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
		},
		AccessToken: token,
		ExpiresIn: 900,
	}, nil
}