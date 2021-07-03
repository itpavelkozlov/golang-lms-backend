package domain

import (
	"context"
	"time"
)

type User struct {
	Email     string     `json:"email" db:"email"`
	ID        string     `json:"id" db:"id"`
	Password  string     `json:"password" db:"password"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	LastName  string     `json:"last_name" db:"lastname"`
	FirstName string     `json:"first_name" db:"firstname"`
}

type UserService interface {
	GetByID(ctx context.Context, id string) (*User, error)
}

type UserRepository interface {
	GetByID(ctx context.Context, id string) (*User, error)
}
