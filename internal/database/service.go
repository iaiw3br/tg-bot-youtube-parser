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

func ExistVideo(db *pgxpool.Pool, videoId string) (bool, error) {
	var exist bool
	query := `
		SELECT count(*) > 0 as exist 
		FROM public.videos 
		WHERE id = $1;`

	err := db.QueryRow(context.Background(), query, videoId).Scan(&exist)
	if err != nil {
		return false, err
	}

	return exist, nil
}
