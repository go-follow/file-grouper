package validation

import "github.com/go-playground/validator/v10"

// Validator интерфейс для валидации данных
type Validator interface {
	// Validate валидация входных данных
	Validate(i interface{}) error
}

// Validation валидация
type Validation struct {
	valid *validator.Validate
}

// New инициализация валидации
func New() *Validation {
	return &Validation{
		valid: validator.New(),
	}
}

// Validate реализация валидации входных данных
func (v *Validation) Validate(s interface{}) error {
	return v.valid.Struct(s)
}
