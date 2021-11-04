package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func ConnectDB() (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), viper.GetString("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
