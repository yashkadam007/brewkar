package service

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/yashkadam007/brewkar/internal/config"
	"github.com/yashkadam007/brewkar/internal/domain"
	"github.com/yashkadam007/brewkar/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(email, password, displayName string) (*domain.User, string, error)
	Login(email, password string) (*domain.User, string, error)
	RefreshToken(refreshToken string) (string, string, error)
}

type authService struct {
	userRepo repository.UserRepository
	jwtCfg   config.JWTConfig
}

func NewAuthService(userRepo repository.UserRepository, jwtCfg config.JWTConfig) AuthService {
	return &authService{
		userRepo: userRepo,
		jwtCfg:   jwtCfg,
	}
}

func (s *authService) Register(email, password, displayName string) (*domain.User, string, error) {
	// Check if user already exists
	existing, err := s.userRepo.GetByEmail(email)
	if err == nil && existing != nil {
		return nil, "", errors.New("user with this email already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}

	// Create user
	user := &domain.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		DisplayName:  displayName,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, "", err
	}

	// Generate token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) Login(email, password string) (*domain.User, string, error) {
	// Find user by email
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	// Update last login
	now := time.Now()
	user.LastLoginAt = &now
	user.UpdatedAt = now
	if err := s.userRepo.Update(user); err != nil {
		return nil, "", err
	}

	// Generate token
	token, err := s.generateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func (s *authService) RefreshToken(refreshToken string) (string, string, error) {
	// TODO: Implement refresh token logic
	return "", "", nil
}

func (s *authService) generateToken(userID uuid.UUID) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"exp": time.Now().Add(time.Duration(s.jwtCfg.AccessTokenExp) * time.Minute).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	return token.SignedString([]byte(s.jwtCfg.Secret))
}
