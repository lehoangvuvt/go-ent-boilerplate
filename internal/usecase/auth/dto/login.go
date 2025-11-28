package authusecasedto

import "github.com/go-playground/validator/v10"

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var validate = validator.New(validator.WithRequiredStructEnabled())

func (req *LoginRequest) Validate() error {
	errs := validate.Var(req.Email, "required,email")
	if errs != nil {
		return errs
	}
	errs = validate.Var(req.Password, "required,min=5")
	if errs != nil {
		return errs
	}
	return nil
}
