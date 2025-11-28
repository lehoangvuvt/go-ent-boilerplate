package authusecase

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	securityports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/security"
	authusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth/dto"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userRepo      repositoryports.UserRepository
	jwtService    securityports.JWTService
	tokenDuration time.Duration
}

func NewLoginUsercase(userRepo repositoryports.UserRepository, jwtService securityports.JWTService, tokenDuration time.Duration) *LoginUsecase {
	return &LoginUsecase{
		userRepo:      userRepo,
		jwtService:    jwtService,
		tokenDuration: tokenDuration,
	}
}

func (uc *LoginUsecase) Execute(ctx context.Context, req *authusecasedto.LoginRequest) (*authusecasedto.LoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if errPassword != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	now := time.Now()
	claims := &authusecasedto.AuthClaims{
		UserID: user.ID.String(),
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(uc.tokenDuration)),
			Subject:   user.ID.String(),
		},
	}
	token, err := uc.jwtService.Generate(claims)
	if err != nil {
		return nil, fmt.Errorf("generate token: %w", err)
	}
	return &authusecasedto.LoginResponse{
		Token: token,
		User: &authusecasedto.UserResponseBrief{
			ID:    user.ID.String(),
			Email: user.Email,
		},
	}, nil
}
