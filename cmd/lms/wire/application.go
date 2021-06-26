package wire

import "github.com/jackc/pgx/v4"

type Application struct {
	conn *pgx.Conn
}

func NewApplication(conn *pgx.Conn) Application {
	return Application{
		conn: conn,
	}
}
