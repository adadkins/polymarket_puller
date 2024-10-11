package polymarket_scraper

// func scrapeHolders
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"polymarket_puller/polymarket_scraper/types"
	"time"
)

func (a *PolyScraper) Scrape() {
	ticker := time.NewTicker(a.timeDelay)
	defer ticker.Stop()

	for range ticker.C {
		a.scrapeHolders("0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917")
	}
}

func (a *PolyScraper) scrapeHolders(market string) {
	holders, err := a.getHolders((market))
	if err != nil {
		return
	}

	for _, sides := range holders {
		for _, holder := range sides.Holders {
			err := a.insertHolder(holder)
			if err != nil {
				log.Default().Println(err)
			}
		}
	}
	log.Println("scraped successful. Len: ", (len(holders[0].Holders) + len(holders[1].Holders)))
}

// 0xdd22472e552920b8438158ea7238bfadfa4f736aa4cee91a6b86c39ead110917
func (a *PolyScraper) getHolders(market string) (types.HoldersResponse, error) {
	now := time.Now()

	baseURL := "https://data-api.polymarket.com/"
	url := fmt.Sprintf("%s%s%s%s", baseURL, "holders?market=", market, "&limit=100")
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("accept-language", "en-US,en;q=0.9")
	req.Header.Add("origin", "https://polymarket.com")
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"129\", \"Not=A?Brand\";v=\"8\", \"Chromium\";v=\"129\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	responseBody := types.HoldersResponse{}
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		return nil, err
	}

	//populate the market and the timestamp
	for i := range responseBody {
		for ii := range responseBody[i].Holders {
			responseBody[i].Holders[ii].MarketID = market
			responseBody[i].Holders[ii].CreatedAt = now
		}
	}
	return responseBody, nil
}

func (a *PolyScraper) insertHolder(holder types.Holder) error {
	query := `
    INSERT INTO holder (
        proxy_wallet, bio, asset, pseudonym, amount, 
        display_username_public, outcome_index, name, 
        profile_image, profile_image_optimized, market_id, created_at
    ) 
    VALUES (:proxy_wallet, :bio, :asset, :pseudonym, :amount, 
            :display_username_public, :outcome_index, :name, 
            :profile_image, :profile_image_optimized, :market_id, :created_at)`

	_, err := a.db.NamedExec(query, holder)
	if err != nil {
		return err
	}

	return nil
}
