package userusecase

import (
	"context"

	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	userusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user/dto"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUsecase struct {
	userRepo repositoryports.UserRepository
}

func NewUserUsecase(userRepo repositoryports.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{userRepo: userRepo}
}

func (uc *CreateUserUsecase) Execute(ctx context.Context, req *userusecasedto.CreateUserRequest) (*userdomain.User, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {

		return nil, err
	}

	du := &userdomain.User{
		Email:          req.Email,
		HashedPassword: string(hashed),
	}

	if err := du.Validate(); err != nil {
		return nil, err
	}
	newUser, err := uc.userRepo.Create(ctx, du)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
