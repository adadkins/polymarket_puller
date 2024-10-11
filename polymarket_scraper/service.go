package polymarket_scraper

import (
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PolyScraper struct {
	db        *sqlx.DB
	baseURL   string
	timeDelay time.Duration
}

func NewScraper(connectionString string, baseURL string, timeDelay int) *PolyScraper {
	db, err := initDB(connectionString)
	if err != nil {
		panic(err)
	}

	service := &PolyScraper{
		db:        db,
		baseURL:   baseURL,
		timeDelay: time.Duration(timeDelay) * time.Second,
	}
	return service
}

func initDB(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Ping....")
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping failed: %w", err)
	}

	log.Println("Pong!")
	return db, nil
}
