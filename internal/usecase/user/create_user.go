package userusecase

import (
	"context"
	"encoding/json"

	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
	mailports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/mail"
	queueports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/queue"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
	userusecasedto "github.com/lehoangvuvt/go-ent-boilerplate/internal/usecase/user/dto"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserUsecase struct {
	userRepo      repositoryports.UserRepository
	mailService   mailports.MailService
	queueProducer queueports.QueueProducer
}

func NewUserUsecase(
	userRepo repositoryports.UserRepository,
	mailService mailports.MailService,
	queueProducer queueports.QueueProducer,
) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo:      userRepo,
		mailService:   mailService,
		queueProducer: queueProducer,
	}
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

	type EmailTask struct {
		Email   string `json:"email"`
		Subject string `json:"subject"`
		Body    string `json:"body"`
	}

	task := EmailTask{
		Email:   newUser.Email,
		Subject: "Register Completed!!!",
		Body:    "Welcome to VPAY! Please click this link to activate your account: https://google.com.vn",
	}

	payload, _ := json.Marshal(task)
	opts := &queueports.PublishOptions{
		Headers: map[string]string{
			"type": "register_email",
		},
	}
	_ = uc.queueProducer.Publish(ctx, "send_email", payload, opts)

	return newUser, nil
}
