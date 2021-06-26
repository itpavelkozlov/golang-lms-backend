package postgres

import (
	"context"
	"github.com/itpavelkozlov/golang-lms-backend/internal/domain"
	"github.com/itpavelkozlov/golang-lms-backend/pkg/logger"
	"github.com/jmoiron/sqlx"
)

type postgresUserRepository struct {
	conn   *sqlx.DB
	logger logger.Logger
}

func (p postgresUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	user := new(domain.User)
	err := p.conn.Get(user, "select * from users where id = $1", id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// NewPostgresUserRepository will create an object that represent the user.Repository interface
func NewPostgresUserRepository(conn *sqlx.DB, logger logger.Logger) domain.UserRepository {
	return &postgresUserRepository{conn, logger}
}
