package service

import (
	"time"

	"github.com/aruncs31s/esdcauthmodule/dto"
	"github.com/aruncs31s/esdcauthmodule/repository"
	"github.com/aruncs31s/esdcauthmodule/utils"
	model "github.com/aruncs31s/esdcmodels"
	userRepo "github.com/aruncs31s/esdcusermodule/repository"
)

type AuthService interface {
	AuthServiceReader
	AuthServiceWriter
	// VerifyEmail(token string) error
}
type authService struct {
	reader AuthServiceReader
	writer AuthServiceWriter
}
type AuthServiceReader interface {
	Login(email, password string) (string, error)
}

type authServiceReader struct {
	authRepo   repository.AuthRepository
	jwtService JWTService
}
type AuthServiceWriter interface {
	Register(user dto.RegisterRequest) error
	// ResetPassword(token, newPassword string) error
	// ForgotPassword(email string) error
}
type authServiceWriter struct {
	// authRepo repository.AuthRepository
	userRepo userRepo.UserRepository
}

func newAuthServiceReader(authRepo repository.AuthRepository, jwtService JWTService) AuthServiceReader {
	return &authServiceReader{
		authRepo:   authRepo,
		jwtService: jwtService,
	}
}

func newAuthServiceWriter(userRepo userRepo.UserRepository) AuthServiceWriter {
	return &authServiceWriter{
		userRepo: userRepo,
	}
}

func NewAuthService(
	authRepo repository.AuthRepository, userRepo userRepo.UserRepository, jwtService JWTService,
) AuthService {
	reader := newAuthServiceReader(authRepo, jwtService)
	writer := newAuthServiceWriter(userRepo)
	return &authService{
		reader: reader,
		writer: writer,
	}
}

func (s *authService) Login(email, password string) (string, error) {
	return s.reader.Login(email, password)
}

func (s *authService) Register(user dto.RegisterRequest) error {
	return s.writer.Register(user)
}

func (s *authServiceReader) Login(email, password string) (string, error) {
	// Check if the user exists
	user, err := s.authRepo.FindByEmail(email)
	if err != nil {
		return "", utils.ErrUserNotExists
	}
	
	if user.Password != password {
		return "", utils.ErrPasswordDoesNotMatch
	}
	// Generate JWT token
	token, err := s.jwtService.CreateToken(user.Username, user.Email, user.Role, user.Name)
	if err != nil {
		return "", utils.ErrGeneratingJWT
	}
	return token, nil
}

func (s *authServiceWriter) Register(user dto.RegisterRequest) error {
	newUser := &model.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     "user", // Default role
		Password: user.Password,
		Github: &model.Github{
			Username: getGithubUsername(user.GithubUsername),
		},
		CreatedAt: time.Time{}.Unix(),
		UpdatedAt: time.Time{}.Unix(),
	}
	err := s.userRepo.CreateUser(newUser)
	if err != nil {
		return err
	}
	return nil
}
func getGithubUsername(username string) string {
	if username == "" {
		return "anonymous"
	}
	return username
}
