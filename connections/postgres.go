package connections

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

type PostgresDBComposite struct {
	DB *pgxpool.Pool
}

func NewPostgresDBComposite() (*PostgresDBComposite, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("DBHOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBPASSWORD"), os.Getenv("DBNAME"))

	database, err := pgxpool.Connect(context.Background(), psqlInfo)
	if err != nil {
		return nil, err
	}

	err = database.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return &PostgresDBComposite{DB: database}, nil
}
