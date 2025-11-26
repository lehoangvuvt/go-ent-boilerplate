package userrepository

import (
	"github.com/lehoangvuvt/go-ent-boilerplate/ent"
	userdomain "github.com/lehoangvuvt/go-ent-boilerplate/internal/domain/user"
)

func toDomain(u *ent.User) *userdomain.User {
	return &userdomain.User{
		ID:             u.ID,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		CreatedAt:      u.CreatedAt,
		UpdateAt:       u.UpdatedAt,
	}
}

func applyDomainToCreate(c *ent.UserCreate, du *userdomain.User) {
	c.SetEmail(du.Email)
	c.SetHashedPassword(du.HashedPassword)
}

func applyDomainToUpdate(uo *ent.UserUpdateOne, du *userdomain.User) {
	uo.SetEmail(du.Email)
	uo.SetHashedPassword(du.HashedPassword)
}
