package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetAllURLs(db *pgxpool.Pool) ([]string, error) {
	query := `SELECT url FROM public.urls;`
	rows, err := db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	var urls []string

	for rows.Next() {
		var url string

		err = rows.Scan(&url)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}
