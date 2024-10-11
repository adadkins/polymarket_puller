package main

import (
	"fmt"
	"os"
	"polymarket_puller/polymarket_scraper"
)

func main() {
	user := os.Getenv("PG_USER")
	dbname := os.Getenv("PG_DB")
	password := os.Getenv("PG_PASSWORD")
	databaseURL := os.Getenv("PG_HOST")
	schema := os.Getenv("PG_SCHEMA")
	connectionString := fmt.Sprintf("host=%s port=5432 user=%s dbname=%s sslmode=disable password=%s search_path=%s", databaseURL, user, dbname, password, schema)
	baseURL := os.Getenv("BASE_URL")
	service := polymarket_scraper.NewScraper(connectionString, baseURL, 120)

	go service.Scrape()

	select {}
}
