package repositoryports

import (
	"context"

	"github.com/google/uuid"
	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
)

type UserRepository interface {
	Create(ctx context.Context, u *userdomain.User) (*userdomain.User, error)
	Update(ctx context.Context, u *userdomain.User) (*userdomain.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	FindByID(ctx context.Context, id uuid.UUID) (*userdomain.User, error)
	FindByEmail(ctx context.Context, email string) (*userdomain.User, error)
}
