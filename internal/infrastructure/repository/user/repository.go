package userrepository

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	"github.com/lehoangvuvt/go-ent-boilerplate/ent/user"
	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
	repositoryports "github.com/lehoangvuvt/go-ent-boilerplate/internal/interface/core/ports/repository"
)

type UserRepository struct {
	client *ent.Client
}

var _ repositoryports.UserRepository = (*UserRepository)(nil)

func NewUserRepository(client *ent.Client) *UserRepository {
	return &UserRepository{
		client: client,
	}
}

func (r *UserRepository) Create(ctx context.Context, du *userdomain.User) (*userdomain.User, error) {
	builder := r.client.User.Create()
	applyDomainToCreate(builder, du)
	nu, err := builder.Save(ctx)
	if ent.IsConstraintError(err) {
		if strings.Contains(err.Error(), "user_email_key") {
			return nil, repositoryports.ErrEmailAlreadyTaken
		}
	}
	if err != nil {
		return nil, err
	}
	return toDomain(nu), nil
}

func (r *UserRepository) Update(ctx context.Context, du *userdomain.User) (*userdomain.User, error) {
	builder := r.client.User.UpdateOneID(du.ID)
	applyDomainToUpdate(builder, du)
	uu, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(uu), nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.client.User.DeleteOneID(id).Exec(ctx)
}

func (r *UserRepository) FindByID(ctx context.Context, id uuid.UUID) (*userdomain.User, error) {
	u, err := r.client.User.Get(ctx, id)
	if ent.IsNotFound(err) {
		return nil, repositoryports.ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*userdomain.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return toDomain(u), nil
}
