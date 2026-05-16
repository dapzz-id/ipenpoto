package services

import (
	"errors"

	"ipenpoto/app/enums"
	"ipenpoto/app/repositories"
	"ipenpoto/app/responses"
	"ipenpoto/database/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo                   *repositories.UserRepository
	JWTService                 *JWTService
	EmailVerificationService   *EmailVerificationService
	EmailService               *EmailService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	jwtService *JWTService,
	emailVerificationService *EmailVerificationService,
	emailService *EmailService,
) *AuthService {
	return &AuthService{
		UserRepo:                 userRepo,
		JWTService:               jwtService,
		EmailVerificationService: emailVerificationService,
		EmailService:             emailService,
	}
}

func (s *AuthService) Login(
	username string,
	password string,
) (
	*responses.LoginResponse,
	string,
	string,
	error,
) {
	const invalidCreds = "Invalid username or password"

	user, err := s.UserRepo.FindByUsername(username)
	if err != nil {
		return nil, "", "", errors.New(invalidCreds)
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)
	if err != nil {
		return nil, "", "", errors.New(invalidCreds)
	}

	if user.EmailVerifiedAt == nil {
		return nil, "", "", errors.New("Email not verified. Check your inbox to verify your email")
	}

	accessToken, err := s.JWTService.GenerateAccessToken(
		user.ID,
		user.Username,
		user.Role,
	)
	if err != nil {
		return nil, "", "", err
	}

	refreshToken, err := s.JWTService.GenerateRefreshToken(
		user.ID,
	)
	if err != nil {
		return nil, "", "", err
	}

	return &responses.LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     string(user.Role),
	}, accessToken, refreshToken, nil
}

func (s *AuthService) Register(
	name string,
	username string,
	email string,
	password string,
) (
	*models.User,
	error,
) {
	existingUser, _ := s.UserRepo.FindByUsername(username)
	if existingUser != nil {
		return nil, errors.New("Username already exists")
	}

	existingEmail, _ := s.UserRepo.FindByEmail(email)
	if existingEmail != nil {
		return nil, errors.New("Email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		ID:       uuid.New(),
		Name:     name,
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
		Role:     enums.Customer,
		Status:   enums.Pending,
	}

	err = s.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}

	verificationToken, err := s.EmailVerificationService.GenerateToken(user.ID)
	if err != nil {
		return nil, errors.New("Failed to generate verification token")
	}

	err = s.EmailService.SendVerificationEmail(user, verificationToken)
	if err != nil {
		return nil, errors.New("Failed to send verification email")
	}

	return user, nil
}

func (s *AuthService) VerifyEmail(token string) error {
	userID, err := s.EmailVerificationService.VerifyToken(token)
	if err != nil {
		return err
	}

	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return errors.New("User not found")
	}

	if user.EmailVerifiedAt != nil {
		return errors.New("Email already verified")
	}

	err = s.UserRepo.VerifyEmail(userID)
	if err != nil {
		return errors.New("Failed to verify email")
	}

	return nil
}
