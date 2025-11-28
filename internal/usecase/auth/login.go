package authusecase

import (
	"context"
	"fmt"

	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	authusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/auth/dto"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userRepo repositoryports.UserRepository
}

func NewLoginUsercase(userRepo repositoryports.UserRepository) *LoginUsecase {
	return &LoginUsecase{
		userRepo: userRepo,
	}
}

func (uc *LoginUsecase) Execute(ctx context.Context, req *authusecasedto.LoginRequest) (*userdomain.User, error) {
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
	return user, nil
}
